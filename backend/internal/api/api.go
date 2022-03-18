package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/api/oapi"
	"github.com/discord-gophers/goapi-gen/pkg/middleware"
	"github.com/discord-gophers/goapi-gen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
)

type ctxKey uint8

const (
	_ ctxKey = iota
	sessionKey
)

// Handler creates a new API handler.
func Handler(server foodtinder.Server) http.Handler {
	handler := handler{server}
	h := oapi.Handler(handler, oapi.WithErrorHandler(handler.onError))

	validator := middleware.OapiRequestValidatorWithOptions(nil, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: handler.authenticate,
		},
	})

	r := chi.NewRouter()
	r.Use(injectSessionBox)
	r.Use(validator)
	r.Mount("/", h)

	return r
}

type handler struct {
	Server foodtinder.Server
}

type authError struct {
	error
}

func (h handler) onError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case oapi.ParameterError:
		respondJSON(w, 400, oapi.FormError{
			FormID: optstr(err.ParamName()),
			Error:  oapi.RespErr(err),
		})
	case authError:
		respondJSON(w, 401, oapi.RespErr(err))
	default:
		respondJSON(w, 400, oapi.RespErr(err))
	}
}

func (h handler) authenticate(ctx context.Context, in *openapi3filter.AuthenticationInput) error {
	switch in.SecuritySchemeName {
	case "BearerAuth":
		auth := in.RequestValidationInput.Request.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return fmt.Errorf("invalid BearerAuth header: missing prefix")
		}

		auth = strings.TrimPrefix(auth, "Bearer ")

		s, err := h.Server.AuthorizerServer().Authorize(ctx, auth)
		if err != nil {
			return authError{err}
		}

		authorizeRequest(in.RequestValidationInput.Request, s)
		return nil
	default:
		return fmt.Errorf("unsupported auth scheme %q", in.SecuritySchemeName)
	}
}

func (h handler) Login(w http.ResponseWriter, r *http.Request, params oapi.LoginParams) *oapi.Response {
	s, err := h.Server.LoginServer().Login(
		r.Context(),
		params.Username, params.Password,
		foodtinder.LoginMetadata{
			UserAgent: r.Header.Get("User-Agent"),
		},
	)
	if err != nil {
		return oapi.LoginJSON401Response(oapi.RespErr(err))
	}

	return oapi.LoginJSON200Response(oapi.Session{
		UserID: oapi.ID(s.UserID),
		Token:  s.Token,
		Expiry: s.Expiry,
		Metadata: oapi.LoginMetadata{
			UserAgent: optstr(s.Metadata.UserAgent),
		},
	})
}

func (h handler) Register(w http.ResponseWriter, r *http.Request, params oapi.RegisterParams) *oapi.Response {
	s, err := h.Server.LoginServer().Register(
		r.Context(),
		params.Username, params.Password,
		foodtinder.LoginMetadata{
			UserAgent: r.Header.Get("User-Agent"),
		},
	)
	if err != nil {
		return oapi.RegisterJSON400Response(oapi.FormError{
			Error: oapi.RespErr(err),
		})
	}

	return oapi.LoginJSON200Response(oapi.Session{
		UserID: oapi.ID(s.UserID),
		Token:  s.Token,
		Expiry: s.Expiry,
		Metadata: oapi.LoginMetadata{
			UserAgent: optstr(s.Metadata.UserAgent),
		},
	})
}

func (h handler) GetSelf(w http.ResponseWriter, r *http.Request) *oapi.Response {
	srv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))

	self, err := srv.Self(r.Context())
	if err != nil {
		return oapi.GetSelfJSON500Response(oapi.RespErr(err))
	}

	return oapi.GetSelfJSON200Response(oapi.Self{
		User: oapi.User{
			Avatar: string(self.Avatar),
			Bio:    optstr(self.Bio),
			ID:     oapi.ID(self.ID),
			Name:   self.Name,
		},
		Birthday: types.Date{
			Time: self.Birthday.Time(),
		},
	})
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request, id oapi.ID, params oapi.GetUserParams) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	usrv := asrv.UserServer()

	u, err := usrv.User(r.Context(), foodtinder.ID(id))
	if err != nil {
		if errors.Is(err, foodtinder.ErrNotFound) {
			return oapi.GetUserJSON404Response(oapi.FormError{
				FormID: optstr("id"),
				Error:  oapi.RespErr(err),
			})
		}
		return oapi.GetUserJSON500Response(oapi.RespErr(err))
	}

	return oapi.GetUserJSON200Response(oapi.User{
		Avatar: string(u.Avatar),
		Bio:    optstr(u.Bio),
		ID:     oapi.ID(u.ID),
		Name:   u.Name,
	})
}

func (h handler) GetNextPosts(w http.ResponseWriter, r *http.Request, params oapi.GetNextPostsParams) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	psrv := asrv.PostServer()

	var prevID foodtinder.ID
	if params.PrevID != nil {
		prevID = foodtinder.ID(*params.PrevID)
	}

	p, err := psrv.NextPosts(r.Context(), prevID)
	if err != nil {
		return oapi.GetNextPostsJSON400Response(oapi.FormError{
			FormID: optstr("prev_id"),
			Error:  oapi.RespErr(err),
		})
	}

	return oapi.GetNextPostsJSON200Response(convertPostsToOAPI(p))
}

func (h handler) DeletePost(w http.ResponseWriter, r *http.Request, id oapi.ID) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	psrv := asrv.PostServer()

	if err := psrv.DeletePost(r.Context(), foodtinder.ID(id)); err != nil {
		if errors.Is(err, foodtinder.ErrNotFound) {
			return oapi.DeletePostJSON404Response(oapi.FormError{
				FormID: optstr("id"),
				Error:  oapi.RespErr(err),
			})
		}
		return oapi.DeletePostJSON500Response(oapi.RespErr(err))
	}

	return nil
}

func (h handler) GetLikedPosts(w http.ResponseWriter, r *http.Request) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	psrv := asrv.PostServer()

	p, err := psrv.LikedPosts(r.Context())
	if err != nil {
		return oapi.GetLikedPostsJSON500Response(oapi.RespErr(err))
	}

	return oapi.GetLikedPostsJSON200Response(convertPostsToOAPI(p))
}

func (h handler) GetAsset(w http.ResponseWriter, r *http.Request, id string) *oapi.Response {
	f, err := h.Server.AssetServer().Open(id)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return oapi.GetAssetJSON404Response(oapi.FormError{
				FormID: optstr("id"),
				Error:  oapi.RespErr(err),
			})
		}
		return oapi.GetAssetJSON500Response(oapi.RespErr(err))
	}
	defer f.Close()

	// No error-handling can be done here.
	io.Copy(w, f)
	return nil
}

func convertPostsToOAPI(posts []foodtinder.Post) []oapi.Post {
	conv := make([]oapi.Post, len(posts))
	for i, p := range posts {
		conv[i] = convertPostToOAPI(p)
	}
	return conv
}

func convertPostToOAPI(post foodtinder.Post) oapi.Post {
	return oapi.Post{
		ID:          oapi.ID(post.ID),
		UserID:      oapi.ID(post.UserID),
		CoverHash:   optstr(post.CoverHash),
		Images:      post.Images,
		Description: post.Description,
		Tags:        post.Tags,
		Location:    optstr(post.Location),
	}
}

func optstr(str string) *string {
	if str != "" {
		return &str
	}
	return nil
}

func respondJSON(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
