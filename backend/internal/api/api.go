package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/acmCSUFDev/Food-Tinder/backend/internal/api/oapi"
	"github.com/discord-gophers/goapi-gen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
)

type ctxKey uint8

const (
	_ ctxKey = iota
	sessionKey
)

type opts struct {
	log *log.Logger
}

func (o *opts) Println(v ...interface{}) {
	if o.log != nil {
		o.log.Println(v...)
	}
}

// Handler creates a new API handler.
func Handler(server foodtinder.Server) http.Handler {
	handler := handler{server}

	r := chi.NewRouter()
	r.NotFound(handler.onNotFound)

	r.Use(injectSessionBox)

	r.Use(oapi.RegisterSecurity(oapi.RegisterSecurityFunc{
		Func:    handler.authenticate,
		OnError: handler.onError,
	}))

	oapi.Handler(
		handler,
		oapi.WithServerBaseURL("/api/v0"),
		oapi.WithErrorHandler(handler.onError),
		oapi.WithRouter(r), // returned http.Handler is useless
	)

	return r
}

type handler struct {
	Server foodtinder.Server
}

type authError struct {
	error
}

func (h handler) onNotFound(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, 404, oapi.RespErr(foodtinder.ErrNotFound))
}

func (h handler) onError(w http.ResponseWriter, r *http.Request, err error) {
	switch err := err.(type) {
	case oapi.ParameterError:
		respondJSON(w, 400, oapi.FormError{
			FormID: err.ParamName(),
			Error:  oapi.RespErr(err),
		})
	case authError:
		respondJSON(w, 401, oapi.RespErr(err))
	case *openapi3filter.SecurityRequirementsError:
		respondJSON(w, 403, oapi.RespErr(securityError(err)))
	default:
		respondJSON(w, 400, oapi.RespErr(err))
	}
}

func securityError(err *openapi3filter.SecurityRequirementsError) error {
	if len(err.Errors) > 0 {
		return err.Errors[0]
	}
	return errors.New("security requirements failed")
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
		Username: s.Username,
		Token:    s.Token,
		Expiry:   s.Expiry,
		Metadata: oapi.LoginMetadata(s.Metadata),
	})
}

func (h handler) Register(w http.ResponseWriter, r *http.Request, params oapi.RegisterParams) *oapi.Response {
	if err := foodtinder.ValidateUsername(params.Username); err != nil {
		return oapi.RegisterJSON400Response(oapi.FormError{
			FormID: "username",
			Error:  oapi.RespErr(err),
		})
	}

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
		Username: s.Username,
		Token:    s.Token,
		Expiry:   s.Expiry,
		Metadata: oapi.LoginMetadata(s.Metadata),
	})
}

func (h handler) GetSelf(w http.ResponseWriter, r *http.Request) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	usrv := asrv.UserServer()

	self, err := usrv.Self(r.Context())
	if err != nil {
		return oapi.GetSelfJSON500Response(oapi.RespErr(err))
	}

	return oapi.GetSelfJSON200Response(oapi.Self{
		User: oapi.User{
			Username:    self.Username,
			DisplayName: self.DisplayName,
			Avatar:      self.Avatar,
			Bio:         self.Bio,
		},
		Birthday: optDate(self.Birthday),
	})
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request, username string) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	usrv := asrv.UserServer()

	u, err := usrv.User(r.Context(), username)
	if err != nil {
		if errors.Is(err, foodtinder.ErrNotFound) {
			return oapi.GetUserJSON404Response(oapi.FormError{
				FormID: "id",
				Error:  oapi.RespErr(err),
			})
		}
		return oapi.GetUserJSON500Response(oapi.RespErr(err))
	}

	return oapi.GetUserJSON200Response(oapi.User{
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Avatar:      u.Avatar,
		Bio:         u.Bio,
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
			FormID: "prev_id",
			Error:  oapi.RespErr(err),
		})
	}

	return oapi.GetNextPostsJSON200Response(convertPostsToOAPI(p))
}

func (h handler) CreatePost(w http.ResponseWriter, r *http.Request) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	usrv := asrv.UserServer()
	psrv := asrv.PostServer()

	var body oapi.Post
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		if errors.Is(err, io.EOF) {
			err = io.ErrUnexpectedEOF
		}
		return oapi.CreatePostJSON400Response(oapi.RespErr(err))
	}

	// Ensure that these are not null, because that's bad.
	if body.Images == nil {
		body.Images = []string{}
	}
	if body.Tags == nil {
		body.Tags = []string{}
	}

	u, err := usrv.Self(r.Context())
	if err != nil {
		return oapi.CreatePostJSON500Response(oapi.RespErr(err))
	}

	// TODO: CoverHash.

	post := foodtinder.Post{
		Images:      body.Images,
		Description: body.Description,
		Tags:        body.Tags,
		Location:    body.Location,
	}

	if err := post.Validate(); err != nil {
		return oapi.CreatePostJSON400Response(oapi.RespErr(err))
	}

	id, err := psrv.CreatePost(r.Context(), post)
	if err != nil {
		return oapi.CreatePostJSON500Response(oapi.RespErr(err))
	}

	// Modify the parsed oapi.Post directly.
	body.ID = oapi.ID(id)
	body.Username = u.Username

	return oapi.CreatePostJSON200Response(body)
}

func (h handler) DeletePost(w http.ResponseWriter, r *http.Request, id oapi.ID) *oapi.Response {
	asrv := h.Server.AuthorizedServer(sessionFromContext(r.Context()))
	psrv := asrv.PostServer()

	if err := psrv.DeletePost(r.Context(), foodtinder.ID(id)); err != nil {
		if errors.Is(err, foodtinder.ErrNotFound) {
			return oapi.DeletePostJSON404Response(oapi.FormError{
				FormID: "id",
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
	f, err := h.Server.FileServer().Open(id)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return oapi.GetAssetJSON404Response(oapi.FormError{
				FormID: "id",
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

func (h handler) UploadAsset(w http.ResponseWriter, r *http.Request) *oapi.Response {
	body := newMaxBytesReader(r.Body, foodtinder.MaxAssetSize)

	usrv := h.Server.FileServer()
	sess := sessionFromContext(r.Context())

	id, err := usrv.Create(sess, body)
	if err != nil {
		if errors.Is(err, errBodyTooLarge) {
			return oapi.UploadAssetJSON413Response(oapi.RespErr(err))
		}
		return oapi.UploadAssetJSON500Response(oapi.RespErr(err))
	}

	return oapi.UploadAssetJSON200Response(id)
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
		Username:    post.Username,
		CoverHash:   post.CoverHash,
		Images:      post.Images,
		Description: post.Description,
		Tags:        post.Tags,
		Location:    post.Location,
	}
}

func emptystr(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func optDate(date foodtinder.Date) *types.Date {
	if date.IsZero() {
		return nil
	}
	return &types.Date{Time: date.Time()}
}

func respondJSON(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
