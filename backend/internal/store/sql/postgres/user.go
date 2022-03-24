package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/store/sql/postgres/query"
	"github.com/pkg/errors"
)

type UserServer AuthorizedServer

func (s *UserServer) User(ctx context.Context, username string) (*foodtinder.User, error) {
	u, err := s.queries.User(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "SQL error")
	}

	return &foodtinder.User{
		Username:    u.Username,
		DisplayName: u.DisplayName.String,
		Avatar:      u.Avatar.String,
		Bio:         u.Bio.String,
	}, nil
}

func (s *UserServer) Self(ctx context.Context) (*foodtinder.Self, error) {
	u, err := s.queries.Self(ctx, s.session.Username)
	if err != nil {
		return nil, errors.Wrap(err, "SQL error")
	}

	return &foodtinder.Self{
		User: foodtinder.User{
			Username:    u.Username,
			DisplayName: u.DisplayName.String,
			Avatar:      u.Avatar.String,
			Bio:         u.Bio.String,
		},
		Birthday: foodtinder.Date{
			Y: uint16(u.Birthday.Time.Year()),
			M: uint8(u.Birthday.Time.Month()),
			D: uint8(u.Birthday.Time.Day()),
		},
	}, nil
}

func (s *UserServer) UpdateSelf(ctx context.Context, new *foodtinder.Self) error {
	foodPrefs, err := json.Marshal(new.Preferences)
	if err != nil {
		return errors.Wrap(err, "error marshaling preferences")
	}

	err = s.queries.UpdateUser(ctx, query.UpdateUserParams{
		Username:    s.session.Username,
		DisplayName: nullstr(new.DisplayName),
		Avatar:      nullstr(new.Avatar),
		Bio:         nullstr(new.Bio),
		Birthday: sql.NullTime{
			Time: new.Birthday.Time(), Valid: new.Birthday == (foodtinder.Date{}),
		},
		FoodPreferences: foodPrefs,
	})
	return errors.Wrap(err, "SQL error")
}

func (s *UserServer) ChangePassword(ctx context.Context, newPassword string) error {
	passhash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	err = s.queries.ChangePassword(ctx, query.ChangePasswordParams{
		Username: s.session.Username,
		Passhash: passhash,
	})
	return errors.Wrap(err, "SQL error")
}
