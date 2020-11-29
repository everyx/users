package users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mholt/binding"
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

type SignupData struct {
	Email    string                 `json:"email"`
	Password string                 `json:"password"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (s *SignupData) Bind(_ *http.Request) error {
	return nil
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginData) Bind(_ *http.Request) error {
	return nil
}

// Fieldmap for the LoginData
func (l *LoginData) FieldMap(_ *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&l.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&l.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	signupData := &SignupData{Metadata: make(map[string]interface{})}

	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		contentType = jsonContentType
	}

	switch contentType {
	case formContentType:
		r.ParseForm()
		for k, v := range r.Form {

			if k == "email" && len(v) == 0 {
				render.Render(w, r, ErrInvalidRequest(fmt.Errorf("email is required")))
				return
			}

			if k == "password" && len(v) == 0 {
				render.Render(w, r, ErrInvalidRequest(fmt.Errorf("password is required")))
				return
			}

			if len(v) == 0 {
				continue
			}

			if k == "email" && len(v) > 0 {
				signupData.Email = v[0]
				continue
			}

			if k == "password" && len(v) > 0 {
				signupData.Password = v[0]
				continue
			}

			if len(v) == 1 {
				signupData.Metadata[k] = v[0]
				continue
			}
			if len(v) > 1 {
				signupData.Metadata[k] = v
			}
		}

	case jsonContentType:
		if err := render.Bind(r, signupData); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}
	if len(signupData.Password) < 8 {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("password length < 8")))
		return
	}

	if !isEmailValid(signupData.Email) {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("invalid email")))
		return
	}

	hashedPassword, err := hashPassword(signupData.Password)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	id, err := a.userStore.New(signupData.Email, hashedPassword, signupData.Metadata)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	confirmationToken := uuid.New().String()
	err = a.userStore.SaveConfirmationToken(id, confirmationToken)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/confirm/%s", confirmationToken))
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = a.userStore.SaveConfirmationTokenSentAt(id, time.Now())
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	switch contentType {
	case formContentType:
		http.Redirect(w, r, "/login?confirmation_sent=true", http.StatusSeeOther)
		return
	case jsonContentType:
		render.Status(r, http.StatusOK)
		return
	}
}

func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	loginData := new(LoginData)

	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		contentType = jsonContentType
	}

	switch contentType {
	case formContentType:
		if errs := binding.Bind(r, loginData); errs != nil {
			render.Render(w, r, ErrInvalidRequest(errs))
			return
		}

	case jsonContentType:
		if err := render.Bind(r, loginData); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}
	if len(loginData.Password) < 8 {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("password length < 8")))
		return
	}

	if !isEmailValid(loginData.Email) {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("invalid email")))
		return
	}

	id, err := a.userStore.UserIDByEmail(loginData.Email)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	fmt.Println("id ", id)

	confirmed, err := a.userStore.IsEmailConfirmed(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	if !confirmed {
		switch contentType {
		case formContentType:
			http.Redirect(w, r, "/login?not_confirmed=true", http.StatusSeeOther)
			return
		case jsonContentType:
			render.JSON(w, r, map[string]interface{}{
				"confirmed": false,
			})
			render.Status(r, http.StatusOK)
			return
		}
		return
	}

	password, err := a.userStore.GetPassword(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginData.Password))
	if err != nil {
		render.Render(w, r, ErrUnauthorized(err))
		return
	}

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	token := uuid.New().String()
	session.Values["id"] = id
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		render.Render(w, r, ErrInternal(fmt.Errorf("session.Save %w", err)))
		return
	}

	email, metadata, err := a.userStore.UserData(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	switch contentType {
	case formContentType:
		http.Redirect(w, r, "/app", http.StatusSeeOther)
		return
	case jsonContentType:
		render.JSON(w, r, map[string]interface{}{
			"id":       id,
			"email":    email,
			"metadata": metadata,
			"token":    token,
		})
		render.Status(r, http.StatusOK)
		return
	}

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
		session, err := a.sessionStore.Get(r, "auth-session")
		if err != nil {
			render.Render(w, r, ErrInternal(err))
			return
		}

		id, ok := session.Values["id"]
		if !ok {
			contentType := r.Header.Get("Content-type")
			if contentType == "" {
				contentType = jsonContentType
			}
			switch contentType {
			case formContentType:
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return

			case jsonContentType:
				render.Render(w, r, ErrUnauthorized(err))
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
		return "", "", nil, err
	}

	id, ok := session.Values["id"]
	if !ok {
		return "", "", nil, fmt.Errorf("user not logged in")
	}

	userID, ok := id.(string)
	if !ok {
		return "", "", nil, fmt.Errorf("user not logged in")
	}

	email, metadata, err := a.userStore.UserData(userID)
	if err != nil {
		return "", "", nil, fmt.Errorf("user not logged in")
	}

	return userID, email, metadata, nil
}

func (a *API) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	id, err := a.userStore.UserIDByConfirmationToken(token)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	err = a.userStore.MarkConfirmed(id, true)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}
	err = a.userStore.DeleteConfirmToken(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	http.Redirect(w, r, "/login?confirmed=true", http.StatusSeeOther)

}

func (a *API) ChangeEmail(id, newEmail string) error {
	if !isEmailValid(newEmail) {
		return fmt.Errorf("email is invalid")
	}

	emailChangeToken := uuid.New().String()
	err := a.userStore.SaveEmailChangeToken(id, newEmail, emailChangeToken)
	if err != nil {
		return err
	}

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/change/%s", emailChangeToken))
	if err != nil {
		return err
	}

	err = a.userStore.SaveEmailChangeTokenSentAt(id, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (a *API) ConfirmEmailChange(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	id, err := a.userStore.UserIDByEmailChangeToken(token)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("invalid or expired token: %w", err)))
		return
	}

	newEmail, err := a.userStore.GetEmailChange(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = a.userStore.UpdateEmail(id, newEmail)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = a.userStore.DeleteEmailChangeToken(id)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	http.Redirect(w, r, "/account?email_changed=true", http.StatusSeeOther)
}

func (a *API) Recovery(email string) error {

	if !isEmailValid(email) {
		return fmt.Errorf("invalid emaild")
	}

	id, err := a.userStore.UserIDByEmail(email)
	if err != nil {
		return err
	}

	recoveryToken := uuid.New().String()
	err = a.userStore.SaveRecoveryToken(id, recoveryToken)
	if err != nil {
		return err
	}

	err = a.sendMail(fmt.Sprintf("http://localhost:4000/confirm/%s", recoveryToken))
	if err != nil {
		return err
	}

	err = a.userStore.SaveRecoveryTokenSentAt(id, time.Now())
	if err != nil {
		return err
	}

	return nil
}

type RecoveryData struct {
	Password string `json:"password"`
}

func (re *RecoveryData) Bind(_ *http.Request) error {
	return nil
}

// Fieldmap for the RecoveryData
func (re *RecoveryData) FieldMap(_ *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&re.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

func (a *API) ConfirmRecovery(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	recoveryData := new(RecoveryData)
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		contentType = jsonContentType
	}

	switch contentType {
	case formContentType:
		if errs := binding.Bind(r, recoveryData); errs != nil {
			render.Render(w, r, ErrInvalidRequest(errs))
			return
		}

	case jsonContentType:
		if err := render.Bind(r, recoveryData); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	}

	if len(recoveryData.Password) < 8 {
		render.Render(w, r, ErrInvalidRequest(fmt.Errorf("password length < 8")))
		return
	}

	id, err := a.userStore.UserIDByRecoveryToken(token)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	hashedPassword, err := hashPassword(recoveryData.Password)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	err = a.userStore.UpdatePassword(id, hashedPassword)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	switch contentType {
	case formContentType:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	case jsonContentType:
		render.Status(r, http.StatusOK)
		return
	}

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
