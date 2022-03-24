package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/pkg/errors"
)

type AuthorizerServer Server

func (s *AuthorizerServer) Authorize(ctx context.Context, token string) (*foodtinder.Session, error) {
	r, err := s.queries.ValidateSession(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No such session, possibly expired. Delete it but don't catch the
			// error, then bail.
			s.queries.DeleteSession(ctx, token)
			return nil, foodtinder.ErrInvalidLogin
		}
		return nil, errors.Wrap(err, "SQL error")
	}

	var metadata foodtinder.LoginMetadata
	json.Unmarshal(r.Metadata, &metadata)

	return &foodtinder.Session{
		Username: r.Username,
		Token:    token,
		Expiry:   r.Expiry,
		Metadata: metadata,
	}, nil
}
