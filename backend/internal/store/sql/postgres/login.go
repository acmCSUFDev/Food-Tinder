package postgres

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres/query"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func hashPassword(password string) ([]byte, error) {
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, errors.Wrap(err, "cannot hash password")
	}
	return passhash, nil
}

type LoginServer Server

func (l *LoginServer) Login(ctx context.Context, username, password string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	passhash, err := l.queries.UserPasshash(ctx, username)
	if err != nil {
		// Assume does not exist.
		return nil, foodtinder.ErrInvalidLogin
	}

	if err := bcrypt.CompareHashAndPassword(passhash, []byte(password)); err != nil {
		// Invalid password.
		return nil, foodtinder.ErrInvalidLogin
	}

	return l.newSession(ctx, username, m)
}

func (l *LoginServer) Register(ctx context.Context, username, password string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	passhash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	err = l.queries.CreateUser(ctx, query.CreateUserParams{
		Username: username,
		Passhash: passhash,
	})
	if err != nil {
		if errorIsUniqueViolation(err) {
			// Unique constraint error, possibly on the username column.
			return nil, foodtinder.ErrUsernameExists
		}
		return nil, errors.Wrap(err, "SQL error")
	}

	return l.newSession(ctx, username, m)
}

func (l *LoginServer) newSession(ctx context.Context, username string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	metadata, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "cannot encode login metadata")
	}

	// Generate the token until we're sure that it doesn't exist. Realistically,
	// this rarely ever happens, but we don't want to scare the user with
	// mundane errors that we're supposed to handle.
	for {
		token := generateToken()

		s, err := l.queries.CreateSession(ctx, query.CreateSessionParams{
			Token:    token,
			Username: username,
			Metadata: metadata,
		})
		if err != nil {
			if errorIsUniqueViolation(err) {
				// A field (probably token) is violating a unique violation.
				// Regenerate a new one.
				continue
			}
			return nil, errors.Wrap(err, "SQL error")
		}

		return &foodtinder.Session{
			Username: username,
			Token:    token,
			Expiry:   s.Expiry,
			Metadata: m,
		}, nil
	}
}

var entropy = base64.URLEncoding.DecodedLen(32)

func generateToken() string {
	b := make([]byte, entropy)

	_, err := rand.Read(b)
	if err != nil {
		log.Panicln("cannot read random:", err)
	}

	return base64.URLEncoding.EncodeToString(b)
}
