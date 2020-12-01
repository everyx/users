package users

import (
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
)

// https://golangcode.com/validate-an-email-address/
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {

	if len(e) < 3 || len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}

	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

type Email struct {
	email.Email
}

type SendmailFunc func(*Email) error

func DefaultSendEmailFunc(cfg Config, pool *email.Pool) func(e *Email) error {
	return func(e *Email) error {
		e.From = cfg.SMTPAdminEmail
		return pool.Send(&e.Email, time.Second*60)
	}
}

func newEmailPool(cfg Config) (*email.Pool, error) {
	var pool *email.Pool
	var err error

	if cfg.SMTPUnencrypted {
		pool, err = email.NewPool(
			fmt.Sprintf("%s:%d", cfg.SMTPHost, cfg.SMTPPort),
			10, &unencryptedAuth{
				smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)},
		)
		if err != nil {
			return nil, err
		}

		return pool, nil
	}

	pool, err = email.NewPool(
		fmt.Sprintf("%s:%d", cfg.SMTPHost, cfg.SMTPPort),
		10,
		smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost),
	)

	if err != nil {
		return nil, err
	}

	return pool, nil
}

type unencryptedAuth struct {
	smtp.Auth
}

// Start starts the auth process for the specified SMTP server.
func (u *unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	server.TLS = true
	return u.Auth.Start(server)
}
