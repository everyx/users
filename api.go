package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

const (
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"
	ctxUserIdKey    = "key_user_id"
)

func NewDefaultAPI(ctx context.Context, driver, dataSource, secret string) (*API, error) {
	userStore, err := NewDefaultUserStore(ctx, driver, dataSource)
	if err != nil {
		return nil, err
	}
	sessionStore, err := NewDefaultSessionStore(ctx, driver, dataSource, []byte(secret))
	if err != nil {
		return nil, err
	}

	return &API{
		ctx:          ctx,
		userStore:    userStore,
		sessionStore: sessionStore,
	}, nil
}

func NewAPI(ctx context.Context, userStore UserStore, sessionStore SessionsStore) *API {
	return &API{
		ctx:          ctx,
		userStore:    userStore,
		sessionStore: sessionStore,
	}
}

type API struct {
	ctx          context.Context
	userStore    UserStore
	sessionStore SessionsStore
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

	id, err := a.userStore.New(email, hashedPassword, metadata)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	confirmationToken := uuid.New().String()
	err = a.userStore.SaveConfirmationToken(id, confirmationToken)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/confirm/%s", confirmationToken))
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

func (a *API) LoggedInUser(r *http.Request) (string, string, map[string]interface{}, error) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return "", "", nil, fmt.Errorf("%v, %w", err, ErrLoginSessionNotFound)
	}

	id, ok := session.Values["id"]
	if !ok {
		return "", "", nil, fmt.Errorf("%w", ErrUserNotLoggedIn)
	}

	userID, ok := id.(string)
	if !ok {
		return "", "", nil, fmt.Errorf("%v %w", err, ErrUserNotLoggedIn)
	}

	email, metadata, err := a.userStore.UserData(userID)
	if err != nil {
		return "", "", nil, fmt.Errorf("%v, %w", err, ErrUserNotLoggedIn)
	}

	return userID, email, metadata, nil
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

	emailChangeToken := uuid.New().String()
	err := a.userStore.SaveEmailChangeToken(id, newEmail, emailChangeToken)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/change/%s", emailChangeToken))
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

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/reset/%s", recoveryToken))
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

func (a *API) UpdateMetaData(id string, metaData map[string]interface{}) error {
	return a.userStore.UpsertMetaData(id, metaData)
}

func (a *API) DeleteMetaDataKeys(id string, keys []string) error {
	return a.userStore.DeleteKeysMetaData(id, keys)
}

func (a *API) sendMail(data string) error {
	log.Println(data)
	return nil
}
