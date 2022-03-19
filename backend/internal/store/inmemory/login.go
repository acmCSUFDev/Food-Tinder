package inmemory

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

// sessionExpiry sets the session to expire after 2 days.
const sessionExpiry = 2 * 24 * time.Hour

type loginServer server

func (s *loginServer) Login(ctx context.Context, username, password string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	s.mu.RLock()
	_, err := s.verifyUser(username, password)
	s.mu.RUnlock()

	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	session, err := s.newSession(ctx, username, m)
	if err != nil {
		return nil, err
	}

	s.sessions[session.Token] = session
	return session, nil
}

func (s *loginServer) verifyUser(username, password string) (*User, error) {
	for i, u := range s.store.Users {
		if u.Username == username && u.Password == password {
			return &s.store.Users[i], nil
		}
	}
	return nil, foodtinder.ErrInvalidLogin
}

func (s *loginServer) Register(ctx context.Context, username, password string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.store.Users {
		if u.Username == username {
			return nil, foodtinder.ErrUsernameExists
		}
	}

	s.store.Users = append(s.store.Users, User{
		Self: foodtinder.Self{
			User: foodtinder.User{
				Username:    username,
				DisplayName: username,
			},
		},
		Password: password,
	})

	session, err := s.newSession(ctx, username, m)
	if err != nil {
		return nil, err
	}

	s.sessions[session.Token] = session
	return session, nil
}

func (s *loginServer) newSession(ctx context.Context, username string, m foodtinder.LoginMetadata) (*foodtinder.Session, error) {
	token, ok := newSessionToken(ctx, func(token string) bool {
		_, exists := s.sessions[token]
		return exists
	})
	if !ok {
		return nil, errors.New("cannot generate a new token")
	}

	now := time.Now()

	return &foodtinder.Session{
		Username: username,
		Token:    token,
		Expiry:   now.Add(sessionExpiry),
		Metadata: m,
	}, nil
}

// entropy is the number of raw bytes to read from /dev/urandom to obtain a
// string of length 32 bytes.
var entropy = base64.URLEncoding.DecodedLen(32)

func newSessionToken(ctx context.Context, exists func(string) bool) (string, bool) {
	b := make([]byte, entropy)
	var s string

	for {
		select {
		case <-ctx.Done():
			return "", false
		default:
		}

		_, err := rand.Read(b)
		if err != nil {
			continue
		}

		s = base64.URLEncoding.EncodeToString(b)
		if !exists(s) {
			return s, true
		}
	}
}
