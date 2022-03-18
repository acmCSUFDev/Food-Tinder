package api

import (
	"context"
	"net/http"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

type session struct {
	*foodtinder.Session
}

func sessionFromContext(ctx context.Context) *foodtinder.Session {
	s, _ := ctx.Value(sessionKey).(*session)
	return s.Session
}

// injectSessionBox exists to inject a &session{} value into all requests,
// including ones that don't have to be authorized. This is a workaround for
// the OpenAPI package's terrible API.
func injectSessionBox(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), sessionKey, &session{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// authorizeRequest authorizes the given request with the given session. This
// method must only be called within routes injected using injectSessionBox.
func authorizeRequest(r *http.Request, s *foodtinder.Session) {
	box := r.Context().Value(sessionKey).(*session)
	box.Session = s
}
