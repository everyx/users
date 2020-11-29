package internal

import (
	"context"
	"encoding/base32"
	"net/http"
	"strings"
	"time"

	"github.com/adnaan/users/internal/models"
	sessionModel "github.com/adnaan/users/internal/models/session"
	gContext "github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// modified form of https://github.com/wader/gormstore/blob/master/gormstore.go

const sessionIDLen = 32

type contextKey string

type DefaultSessionStore struct {
	Driver      string
	DataSource  string
	Ctx         context.Context
	Client      *models.Client
	Codecs      []securecookie.Codec
	SessionOpts *sessions.Options
}

func (ds *DefaultSessionStore) Close() error {
	return ds.Client.Close()
}

// Get returns a session for the given name after adding it to the registry.
func (ds *DefaultSessionStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(ds, name)
}

func (ds *DefaultSessionStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(ds, name)
	opts := *ds.SessionOpts
	session.Options = &opts

	ds.MaxAge(ds.SessionOpts.MaxAge)
	cookie, err := r.Cookie(name)
	if err != nil {
		return session, nil
	}

	if err := securecookie.DecodeMulti(name, cookie.Value, &session.ID, ds.Codecs...); err != nil {
		return session, nil
	}

	s, err := ds.Client.Session.Query().Where(sessionModel.IDEQ(session.ID), sessionModel.ExpiresAtGT(time.Now())).Only(ds.Ctx)
	if err != nil {
		return session, nil
	}

	if err := securecookie.DecodeMulti(session.Name(), s.Data, &session.Values, ds.Codecs...); err != nil {
		return session, nil
	}

	gContext.Set(r, contextKey(name), s)

	return session, nil
}

func (ds *DefaultSessionStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	sess, _ := gContext.Get(r, contextKey(session.Name())).(*models.Session)

	// delete if max age is < 0
	if session.Options.MaxAge < 0 {
		if sess != nil {
			if err := ds.Client.Session.DeleteOne(sess).Exec(ds.Ctx); err != nil {
				return err
			}
		}
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return nil
	}

	data, err := securecookie.EncodeMulti(session.Name(), session.Values, ds.Codecs...)
	if err != nil {
		return err
	}
	now := time.Now()
	expire := now.Add(time.Second * time.Duration(session.Options.MaxAge))

	if sess == nil {
		// generate random session ID key suitable for storage in the db
		session.ID = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(sessionIDLen)), "=")

		newSess, err := ds.Client.Session.Create().SetID(session.ID).SetData(data).SetUpdatedAt(now).SetExpiresAt(expire).Save(ds.Ctx)
		if err != nil {
			return err
		}
		gContext.Set(r, contextKey(session.Name()), newSess)
	} else {
		_, err = ds.Client.Session.UpdateOne(sess).SetData(data).SetUpdatedAt(now).SetExpiresAt(expire).Save(ds.Ctx)
		if err != nil {
			return err
		}
	}

	// set session id cookie
	id, err := securecookie.EncodeMulti(session.Name(), session.ID, ds.Codecs...)
	if err != nil {
		return err
	}
	http.SetCookie(w, sessions.NewCookie(session.Name(), id, session.Options))

	return nil
}

func (ds *DefaultSessionStore) MaxAge(age int) {
	ds.SessionOpts.MaxAge = age
	for _, codec := range ds.Codecs {
		if sc, ok := codec.(*securecookie.SecureCookie); ok {
			sc.MaxAge(age)
		}
	}
}

// MaxLength restricts the maximum length of new sessions to l.
// If l is 0 there is no limit to the size of a session, use with caution.
// The default is 4096 (default for securecookie)
func (ds *DefaultSessionStore) MaxLength(l int) {
	for _, c := range ds.Codecs {
		if codec, ok := c.(*securecookie.SecureCookie); ok {
			codec.MaxLength(l)
		}
	}
}

// Cleanup deletes expired sessions
func (ds *DefaultSessionStore) Cleanup() {
	ds.Client.Session.Delete().Where(sessionModel.ExpiresAtLTE(time.Now())).Exec(ds.Ctx)
}

// PeriodicCleanup runs Cleanup every interval. Close quit channel to stop.
func (ds *DefaultSessionStore) PeriodicCleanup(interval time.Duration, quit <-chan struct{}) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			ds.Cleanup()
		case <-quit:
			return
		}
	}
}
