package postgres

import (
	"context"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/pkg/errors"
)

type AuthorizedServer struct {
	*Server
	session *foodtinder.Session
}

// Logout invalidates the authorizing token.
func (s *AuthorizedServer) Logout(ctx context.Context) error {
	if err := s.queries.DeleteSession(ctx, s.session.Token); err != nil {
		return errors.Wrap(err, "SQL error")
	}
	return nil
}

func (s *AuthorizedServer) PostServer() foodtinder.PostServer {
	return (*PostServer)(s)
}

func (s *AuthorizedServer) UserServer() foodtinder.UserServer {
	return (*UserServer)(s)
}

func (s *AuthorizedServer) Session() *foodtinder.Session {
	return s.session
}
