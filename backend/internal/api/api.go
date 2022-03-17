package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/api/openapi"
	"github.com/discord-gophers/goapi-gen/pkg/types"
)

// Handler creates a new API handler.
func Handler(server foodtinder.Server) http.Handler {
	handler := handler{server}
	return openapi.Handler(
		handler,
		openapi.WithMiddleware("mustAuthorize", handler.mustAuthorize),
		openapi.WithErrorHandler(handler.onError),
	)
}

type ctxKey uint8

const (
	_ ctxKey = iota
	sessionKey
)

// SessionFromContext gets the session struct from the given context.
func SessionFromContext(ctx context.Context) *foodtinder.Session {
	s, _ := ctx.Value(sessionKey).(*foodtinder.Session)
	return s
}

type handler struct {
	server foodtinder.Server
}

func (h handler) onError(w http.ResponseWriter, r *http.Request, err error) {
	h.writeError(w, 400, err)
}

func (h handler) writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(openapi.Error{
		Message: err.Error(),
	})
}

func (h handler) mustAuthorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		s, err := h.server.AuthorizerServer().Authorize(r.Context(), token)
		if err != nil {
			h.onError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), sessionKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h handler) Login(w http.ResponseWriter, r *http.Request, params openapi.LoginParams) {
	s, err := h.server.LoginServer().Login(
		r.Context(),
		params.Username, params.Password,
		foodtinder.LoginMetadata{
			UserAgent: r.Header.Get("User-Agent"),
		},
	)
	if err != nil {
		h.writeError(w, 401, err)
		return
	}

	openapi.LoginJSON200Response(openapi.Session{
		UserID: openapi.ID(s.UserID),
		Token:  s.Token,
		Expiry: s.Expiry,
		Metadata: openapi.LoginMetadata{
			UserAgent: optstr(s.Metadata.UserAgent),
		},
	})
}

func (h handler) Register(w http.ResponseWriter, r *http.Request, params openapi.RegisterParams) {
	s, err := h.server.LoginServer().Register(
		r.Context(),
		params.Username, params.Password,
		foodtinder.LoginMetadata{
			UserAgent: r.Header.Get("User-Agent"),
		},
	)
	if err != nil {
		h.writeError(w, 400, err)
		return
	}

	openapi.LoginJSON200Response(openapi.Session{
		UserID: openapi.ID(s.UserID),
		Token:  s.Token,
		Expiry: s.Expiry,
		Metadata: openapi.LoginMetadata{
			UserAgent: optstr(s.Metadata.UserAgent),
		},
	})
}

func (h handler) GetUsersSelf(w http.ResponseWriter, r *http.Request) {
	srv := h.server.AuthorizedServer(SessionFromContext(r.Context()))

	self, err := srv.Self(r.Context())
	if err != nil {
		h.writeError(w, 500, err)
	}

	openapi.GetUsersSelfJSON200Response(openapi.Self{
		User: openapi.User{
			Avatar: string(self.Avatar),
			Bio:    optstr(self.Bio),
			ID:     openapi.ID(self.ID),
			Name:   self.Name,
		},
		Birthday: types.Date{
			Time: self.Birthday.Time(),
		},
	})
}

func (h handler) GetUsersID(w http.ResponseWriter, r *http.Request, id openapi.ID, params openapi.GetUsersIDParams) {
	asrv := h.server.AuthorizedServer(SessionFromContext(r.Context()))
	usrv := asrv.UserServer()

	u, err := usrv.User(r.Context(), foodtinder.ID(id))
	if err != nil {
		h.writeError(w, 500, err)
	}

	openapi.GetUsersIDJSON200Response(openapi.User{
		Avatar: string(u.Avatar),
		Bio:    optstr(u.Bio),
		ID:     openapi.ID(u.ID),
		Name:   u.Name,
	})
}

func optstr(str string) *string {
	if str != "" {
		return &str
	}
	return nil
}
