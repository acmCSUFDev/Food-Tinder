package inmemory

import (
	"context"
	"errors"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

// sessionRenew sets the duration to renew the session.
const sessionRenew = time.Hour

type authorizerServer server

func (s *authorizerServer) Authorize(ctx context.Context, token string) (*foodtinder.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[token]
	if !ok {
		return nil, errors.New("unknown token")
	}

	s.bumpSession(session)
	return session, nil
}

func (s *authorizerServer) bumpSession(session *foodtinder.Session) {
	now := time.Now()

	lastBumped := session.Expiry.Add(-sessionExpiry)
	// Check by adding the renewal duration. If the time instant is before (less
	// than) now, then it means we're past the don't-renew timeframe, so we
	// renew.
	if lastBumped.Add(sessionRenew).Before(now) {
		go func() {
			s.mu.Lock()
			session.Expiry = time.Now().Add(sessionExpiry)
			s.mu.Unlock()
		}()
	}
}
