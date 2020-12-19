package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
	ctxUserIdKey    = "key_user_id"
)

type Config struct {
	Driver          string
	Datasource      string
	SessionSecret   string
	SendMail        SendMailFunc
	APIMasterSecret string
	GothProviders   []goth.Provider
}

type User struct {
	ID            string                 `json:"id"`
	BillingID     string                 `json:"billing_id"`
	Email         string                 `json:"email"`
	IsAPITokenSet bool                   `json:"is_api_token_set"`
	Metadata      map[string]interface{} `json:"metadata"`
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

	if len(cfg.GothProviders) > 0 {
		goth.UseProviders(cfg.GothProviders...)
		gothic.Store = sessionStore
	}

	return &API{
		ctx:          ctx,
		userStore:    userStore,
		sessionStore: sessionStore,
		sendMail:     cfg.SendMail,
		branca:       branca.NewBranca(cfg.APIMasterSecret),
	}, nil
}

func NewAPI(ctx context.Context, apiMasterSecret string, userStore UserStore, sessionStore SessionsStore, sendMail SendMailFunc) *API {
	return &API{
		ctx:          ctx,
		userStore:    userStore,
		sessionStore: sessionStore,
		sendMail:     sendMail,
		branca:       branca.NewBranca(apiMasterSecret),
	}
}

type API struct {
	ctx          context.Context
	userStore    UserStore
	sessionStore SessionsStore
	sendMail     SendMailFunc
	branca       *branca.Branca
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}

func (a *API) Signup(email, password string, metadata map[string]interface{}) error {

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

	id, err := a.userStore.New(email, hashedPassword, "email_signup", metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	confirmationToken := uuid.New().String()
	err = a.userStore.SaveConfirmationToken(id, confirmationToken)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(Confirmation, confirmationToken, email, metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrSendingEmail)
	}

	err = a.userStore.SaveConfirmationTokenSentAt(id, time.Now())
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
			ctx = context.WithValue(ctx, ctxUserIdKey, id)
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
			ctx = context.WithValue(ctx, ctxUserIdKey, id)
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

	user := &User{
		ID:            userID,
		BillingID:     billingID,
		Email:         email,
		IsAPITokenSet: isAPIKeySet,
		Metadata:      metadata,
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

func (a *API) HandleGothCallback(w http.ResponseWriter, r *http.Request) error {
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

	metadata := map[string]interface{}{
		"name": usr.Name,
	}

	userMap := structs.Map(usr)

	for k, v := range userMap {
		metadata[k] = v
	}

	var id string

	id, err = a.userStore.UserIDByEmail(usr.Email)
	if err != nil {
		// user not found, create a user
		id, err = a.userStore.New(usr.Email, usr.AccessToken, usr.Provider, metadata)
		if err != nil {
			return err
		}
		err = a.userStore.MarkConfirmed(id, true)
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
		log.Println("getSessionVal: ", err)
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
