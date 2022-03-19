package inmemory

import (
	"context"
	"errors"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

type userServer authorizedServer

func (s *userServer) User(ctx context.Context, username string) (*foodtinder.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.store.Users {
		if u.Username == username {
			user := u.User // copy
			return &user, nil
		}
	}

	return nil, foodtinder.ErrNotFound
}

func (s *userServer) Self(ctx context.Context) (*foodtinder.Self, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, err := s.self()
	if err != nil {
		return nil, err
	}

	self := u.Self // copy
	return &self, nil
}

func (s *userServer) FoodPreferences(ctx context.Context) (*foodtinder.FoodPreferences, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, err := s.self()
	if err != nil {
		return nil, err
	}

	prefs := u.Preferences // copy
	return &prefs, nil
}

func (s *userServer) SetFoodPreferences(ctx context.Context, prefs *foodtinder.FoodPreferences) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, err := s.self()
	if err != nil {
		return err
	}

	u.Preferences = *prefs
	return nil
}

func (s *userServer) self() (*User, error) {
	for i, u := range s.store.Users {
		if u.Username == s.session.Username {
			return &s.store.Users[i], nil
		}
	}
	return nil, errors.New("missing current user in database")
}
