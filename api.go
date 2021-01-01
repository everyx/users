package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zpatrick/rbac"

	"github.com/hako/branca"

	"github.com/lithammer/shortuuid/v3"

	"github.com/fatih/structs"

	"github.com/markbates/goth/gothic"

	"github.com/markbates/goth"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

const (
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"
	CtxUserIdKey    = "key_user_id"
)

type Permission struct {
	ActionMatcher rbac.Matcher
	TargetMatcher func(userID string) rbac.Matcher
}

func NewPermission(action string, targetMatcher func(userID string) rbac.Matcher) Permission {
	return Permission{
		ActionMatcher: rbac.GlobMatch(action),
		TargetMatcher: targetMatcher,
	}
}

type Config struct {
	Driver          string
	Datasource      string
	SessionSecret   string
	SendMail        SendMailFunc
	APIMasterSecret string
	GothProviders   []goth.Provider
	Roles           map[string][]Permission
}

type User struct {
	ID            string                 `json:"id"`
	BillingID     string                 `json:"billing_id"`
	Email         string                 `json:"email"`
	IsAPITokenSet bool                   `json:"is_api_token_set"`
	Metadata      map[string]interface{} `json:"metadata"`
	Workspaces    []Workspace            `json:"workspaces"`
}

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func NewDefaultAPI(ctx context.Context, cfg Config) (*API, error) {
	userStore, err := NewDefaultUserStore(ctx, cfg.Driver, cfg.Datasource)
	if err != nil {
		return nil, err
	}
	sessionStore, err := NewDefaultSessionStore(ctx, cfg.Driver, cfg.Datasource, []byte(cfg.SessionSecret))
	if err != nil {
		return nil, err
	}

	workspaceStore, err := NewDefaultWorkspaceStore(ctx, cfg.Driver, cfg.Datasource)
	if err != nil {
		return nil, err
	}

	if len(cfg.GothProviders) > 0 {
		goth.UseProviders(cfg.GothProviders...)
		gothic.Store = sessionStore
	}

	return &API{
		ctx:            ctx,
		userStore:      userStore,
		sessionStore:   sessionStore,
		workspaceStore: workspaceStore,
		sendMail:       cfg.SendMail,
		branca:         branca.NewBranca(cfg.APIMasterSecret),
		roles:          cfg.Roles,
	}, nil
}

func NewAPI(ctx context.Context, apiMasterSecret string,
	userStore UserStore, sessionStore SessionsStore, workspaceStore WorkspaceStore,
	sendMail SendMailFunc, roles map[string][]Permission) *API {
	return &API{
		ctx:            ctx,
		userStore:      userStore,
		sessionStore:   sessionStore,
		workspaceStore: workspaceStore,
		sendMail:       sendMail,
		branca:         branca.NewBranca(apiMasterSecret),
		roles:          roles,
	}
}

type API struct {
	ctx            context.Context
	userStore      UserStore
	sessionStore   SessionsStore
	workspaceStore WorkspaceStore
	sendMail       SendMailFunc
	branca         *branca.Branca
	roles          map[string][]Permission
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}

func (a *API) Signup(email, password, role string, metadata map[string]interface{}) error {

	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	_, err = a.userStore.UserIDByEmail(email)
	if err == nil {
		return fmt.Errorf("%w", ErrUserExists)
	}

	sendMailFunc := func(token, email string) error {
		return a.sendMail(Confirmation, token, email, metadata)
	}

	_, err = a.userStore.New(email, hashedPassword, role, "email_signup", metadata, sendMailFunc)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) Login(w http.ResponseWriter, r *http.Request, email, password string) error {

	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	id, err := a.userStore.UserIDByEmail(email)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	confirmed, err := a.userStore.IsEmailConfirmed(id)
	if err != nil {
		return err
	}

	if !confirmed {
		return fmt.Errorf("%w", ErrEmailNotConfirmed)
	}

	existingPassword, err := a.userStore.GetPassword(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrWrongPassword)
	}

	// set provider
	err = a.userStore.UpdateProvider(id, "email_signup")
	if err != nil {
		return err
	}

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	token := uuid.New().String()
	session.Values["id"] = id
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	session.Values = nil
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *API) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-type")
		if contentType == "" {
			contentType = formContentType
		}

		authHeader := r.Header.Get("Authorization")

		if contentType == jsonContentType && strings.Contains(authHeader, "Bearer") {
			parts := strings.Split(authHeader, "Bearer ")
			if len(parts) != 2 {
				http.Error(w, "Unauthorized: invalid header", 401)
				return
			}

			apiKey, err := a.branca.DecodeToString(parts[1])
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Unauthorized: invalid bearer token", 401)
				return
			}

			id, err := a.userStore.UserIDByAPIKey(apiKey)
			if id == "" || err != nil {
				fmt.Println(err)
				http.Error(w, "Unauthorized: user not found", 401)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUserIdKey, id)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		redirectTo := fmt.Sprintf("/login?from=%s", r.URL.Path)

		session, err := a.sessionStore.Get(r, "auth-session")
		if err != nil {
			switch contentType {
			case formContentType:
				http.Redirect(w, r, redirectTo, http.StatusSeeOther)
				return

			case jsonContentType:
				http.Error(w, "Unauthorized", 401)
				return
			}
			return
		}

		id, ok := session.Values["id"]
		if !ok {
			switch contentType {
			case formContentType:
				http.Redirect(w, r, redirectTo, http.StatusSeeOther)
				return

			case jsonContentType:
				http.Error(w, "Unauthorized", 401)
				return
			}
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUserIdKey, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (a *API) LoggedInUser(r *http.Request) (*User, error) {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return nil, err
	}

	email, billingID, apiKey, metadata, err := a.userStore.UserData(userID)
	if err != nil {
		return nil, fmt.Errorf("%v, %w", err, ErrUserNotLoggedIn)
	}

	isAPIKeySet := false
	if apiKey != "" {
		isAPIKeySet = true
	}

	workspaceMap, err := a.workspaceStore.GetUserWorkspaces(userID)
	if err != nil {
		return nil, fmt.Errorf("%v, %w", err, ErrInternal)
	}

	var workspaces []Workspace
	for id, data := range workspaceMap {
		workspaces = append(workspaces, Workspace{
			ID:   id,
			Name: data[0],
			Role: data[1],
		})
	}

	user := &User{
		ID:            userID,
		BillingID:     billingID,
		Email:         email,
		IsAPITokenSet: isAPIKeySet,
		Metadata:      metadata,
		Workspaces:    workspaces,
	}

	return user, nil
}

func (a *API) ConfirmEmail(token string) error {
	id, err := a.userStore.UserIDByConfirmationToken(token)
	if err != nil {
		return fmt.Errorf("%v, %w", err, ErrUserNotFound)
	}

	err = a.userStore.MarkConfirmed(id, true)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	err = a.userStore.DeleteConfirmToken(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) ChangeEmail(id, newEmail string) error {
	if !isEmailValid(newEmail) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	email, _, _, metadata, err := a.userStore.UserData(id)
	if err != nil {
		return fmt.Errorf("%w", ErrInternal)
	}

	emailChangeToken := uuid.New().String()
	err = a.userStore.SaveEmailChangeToken(id, newEmail, emailChangeToken)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(ChangeEmail, emailChangeToken, email, metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.SaveEmailChangeTokenSentAt(id, time.Now())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) ConfirmEmailChange(token string) error {
	id, err := a.userStore.UserIDByEmailChangeToken(token)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	newEmail, err := a.userStore.GetEmailChange(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.UpdateEmail(id, newEmail)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.DeleteEmailChangeToken(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) Recovery(email string) error {

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	id, err := a.userStore.UserIDByEmail(email)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUserNotFound)
	}

	recoveryToken := uuid.New().String()
	err = a.userStore.SaveRecoveryToken(id, recoveryToken)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	_, _, _, metadata, err := a.userStore.UserData(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(Recovery, recoveryToken, email, metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.SaveRecoveryTokenSentAt(id, time.Now())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) ConfirmRecovery(token, password string) error {

	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	id, err := a.userStore.UserIDByRecoveryToken(token)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUserNotFound)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.UpdatePassword(id, hashedPassword)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) OTP(email string) error {

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	id, err := a.userStore.UserIDByEmail(email)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUserNotFound)
	}

	otp := uuid.New().String()
	err = a.userStore.SaveOTP(id, otp)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	_, _, _, metadata, err := a.userStore.UserData(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(OTP, otp, email, metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.SaveOTPSentAt(id, time.Now())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) ResetAPIToken(r *http.Request) (string, error) {

	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return "", err
	}
	apiKey := shortuuid.New()

	token, err := a.branca.EncodeToString(apiKey)
	if err != nil {
		return "", fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.UpdateAPIKey(userID, apiKey)
	if err != nil {
		return "", fmt.Errorf("%v %w", err, ErrInternal)
	}

	return token, nil
}

func (a *API) LoginWithOTP(w http.ResponseWriter, r *http.Request, otp string) error {
	id, err := a.userStore.UserIDByOTP(otp)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	token := uuid.New().String()
	session.Values["id"] = id
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.userStore.DeleteOTP(id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) UpdateMetaData(r *http.Request, metaData map[string]interface{}) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.UpsertMetaData(userID, metaData)
}

func (a *API) DeleteMetaDataKeys(r *http.Request, keys []string) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.DeleteKeysMetaData(userID, keys)
}

func (a *API) HandleGothCallback(w http.ResponseWriter, r *http.Request, role string, metadata map[string]interface{}) error {
	usr, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	type User struct {
		RawData           map[string]interface{}
		Provider          string
		Email             string
		Name              string
		FirstName         string
		LastName          string
		NickName          string
		Description       string
		UserID            string
		AvatarURL         string
		Location          string
		AccessToken       string
		AccessTokenSecret string
		RefreshToken      string
		ExpiresAt         time.Time
		IDToken           string
	}

	if metadata == nil {
		metadata = map[string]interface{}{
			"name": usr.Name,
		}
	} else {
		metadata["name"] = usr.Name
	}

	userMap := structs.Map(usr)

	for k, v := range userMap {
		metadata[k] = v
	}

	var id string

	id, err = a.userStore.UserIDByEmail(usr.Email)
	if err != nil {
		// user not found, create a user
		id, err = a.userStore.New(usr.Email, usr.AccessToken, role, usr.Provider, metadata, nil)
		if err != nil {
			return err
		}
	} else {
		// otherwise just update provider & metadata
		err = a.userStore.UpdateProvider(id, usr.Provider)
		if err != nil {
			return err
		}
		err = a.userStore.UpsertMetaData(id, metadata)
		if err != nil {
			return err
		}

	}

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	token := uuid.New().String()
	session.Values["id"] = id
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) HandleGothLogin(w http.ResponseWriter, r *http.Request) error {
	// try to get the user without re-authenticating
	if usr, err := gothic.CompleteUserAuth(w, r); err == nil {
		id, err := a.userStore.UserIDByEmail(usr.Email)
		if err != nil {
			return fmt.Errorf("err: %v, %w", err, ErrUserNotFound)
		}
		session, err := a.sessionStore.Get(r, "auth-session")
		if err != nil {
			return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
		}

		token := uuid.New().String()
		session.Values["id"] = id
		session.Values["token"] = token
		err = session.Save(r, w)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
	} else {
		gothic.BeginAuthHandler(w, r)
	}
	return nil
}

func (a *API) HandleGothLogout(w http.ResponseWriter, r *http.Request) error {
	a.Logout(w, r)
	return gothic.Logout(w, r)
}

func (a *API) DeleteUser(r *http.Request) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.DeleteUser(userID)
}

func (a *API) GetSessionVal(r *http.Request, key string) (interface{}, error) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return nil, fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	val, ok := session.Values[key]
	if !ok {
		return nil, fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	return val, nil
}

func (a *API) GetSessionStringVal(r *http.Request, key string) *string {
	val, err := a.GetSessionVal(r, key)
	if err != nil {
		return nil
	}
	strVal, ok := val.(string)
	if !ok {
		log.Printf("err %v is not string", val)
		return nil
	}

	return &strVal
}

func (a *API) SetSessionVal(r *http.Request, w http.ResponseWriter, key string, val interface{}) error {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values[key] = val
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) DelSessionVal(r *http.Request, w http.ResponseWriter, key string) error {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values[key] = nil
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) getUserIDFromSession(r *http.Request) (string, error) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return "", fmt.Errorf("%v, %w", err, ErrLoginSessionNotFound)
	}

	id, ok := session.Values["id"]
	if !ok {
		return "", fmt.Errorf("%w", ErrUserNotLoggedIn)
	}

	userID, ok := id.(string)
	if !ok {
		return "", fmt.Errorf("%v %w", err, ErrUserNotLoggedIn)
	}
	return userID, nil
}

func (a *API) UpdateBillingID(r *http.Request, billingID string) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.UpdateBillingID(userID, billingID)
}

func (a *API) AddRole(r *http.Request, role string) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.AddRole(userID, role)
}

func (a *API) DeleteRole(r *http.Request, role string) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.DeleteRole(userID, role)
}

func (a *API) ClearRoles(r *http.Request) error {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return err
	}
	return a.userStore.ClearRoles(userID)
}

func (a *API) GetRoles(r *http.Request) ([]string, error) {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return nil, err
	}
	return a.userStore.GetRoles(userID)
}

func (a *API) Can(r *http.Request, action, target string) (bool, error) {
	userID, err := a.getUserIDFromSession(r)
	if err != nil {
		return false, err
	}

	userRoles, err := a.userStore.GetRoles(userID)
	if err != nil {
		return false, err
	}

	if len(userRoles) == 0 {
		return false, fmt.Errorf("user has no roles")
	}

	for _, userRole := range userRoles {
		if strings.Contains(userRole, "::") {
			parts := strings.SplitAfter(userRole, "::")
			if len(parts) == 2 {
				// e.g. team::userid
				if parts[1] != userID {
					continue
				}
				userRole = parts[0]
			}
		}
		role, ok := a.roles[userRole]
		if !ok {
			continue
		}

		var rbacPermissions []rbac.Permission
		for _, perm := range role {
			rbacPermissions = append(rbacPermissions, rbac.NewPermission(perm.ActionMatcher, perm.TargetMatcher(userID)))
		}
		rbacRole := rbac.Role{
			RoleID:      userRole,
			Permissions: rbacPermissions,
		}
		can, err := rbacRole.Can(action, target)
		if err != nil {
			log.Printf("can %v %v, err: %v", action, target, err)
			continue
		}
		if can {
			return true, nil
		}

	}

	return false, nil
}

func (a *API) GetWorkspaceRole(r *http.Request, workspaceID string) (string, error) {
	return "", nil
}

func (a *API) InviteMembers(r *http.Request, workspaceID string, members []string) error {

	// is user exist by email id, add user to workspace and send notification
	for _, member := range members {
		uid, err := a.userStore.UserIDByEmail(member)
		if err != nil {
			// user not found. create invitation record to be processed on signup.
			err = a.workspaceStore.CreateInvitation(workspaceID, member, "member")
			if err != nil {
				log.Printf("CreateInvitation err: %+v", err)
				continue
			}

		} else {
			err = a.workspaceStore.WorkspaceUpsertUser(workspaceID, uid, "member")
			if err != nil {
				log.Printf("WorkspaceUpsertUser err: %+v", err)
				continue
			}
		}

		name, _, _, err := a.workspaceStore.GetWorkspace(workspaceID)
		if err != nil {
			log.Printf("workspaceStore.GetWorkspace err: %+v", err)
			continue
		}
		// send invitation email
		metadata := map[string]interface{}{
			"workspace": name,
		}
		err = a.sendMail(InviteMember, "", member, metadata)
		if err != nil {
			log.Printf("sendMail(InviteMember) err: %+v", err)
			continue
		}
	}

	// is user does not exist, send invitation email
	return nil
}
