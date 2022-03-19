package inmemory

import (
	"context"
	"errors"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

type authorizedServer struct {
	*server
	session *foodtinder.Session
}

func (s *authorizedServer) Logout(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.sessions[s.session.Token]; !ok {
		return errors.New("session is already removed")
	}

	delete(s.sessions, s.session.Token)
	return nil
}

func (s *authorizedServer) PostServer() foodtinder.PostServer {
	return (*postServer)(s)
}

func (s *authorizedServer) UserServer() foodtinder.UserServer {
	return (*userServer)(s)
}

func (s *authorizedServer) AssetUploadServer() foodtinder.AssetUploadServer {
	return (*assetUploadServer)(s)
}
