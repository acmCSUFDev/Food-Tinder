package oapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/bwmarrin/snowflake"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

//go:generate goapi-gen --config=config.json ../../../../openapi/foodtinder.jsonc

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(foodtinder.ID(id).String()))
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	s, err := snowflake.ParseString(str)
	if err != nil {
		return fmt.Errorf("invalid snowflake: %v", err)
	}
	*id = ID(s)
	return nil
}

// RespErr returns an Error response value from an error.
func RespErr(err error) Error {
	return Error{err.Error()}
}

// MustSwagger calls GetSwagger or panics.
func MustSwagger() *openapi3.T {
	t, err := GetSwagger()
	if err != nil {
		panic("GetSwagger: " + err.Error())
	}
	return t
}

// RegisterSecurityFunc is a structure containing 2 callbacks needed for
// RegisterSecurity.
type RegisterSecurityFunc struct {
	Func    openapi3filter.AuthenticationFunc
	OnError func(w http.ResponseWriter, r *http.Request, err error)
}

// RegisterSecurity registers the given security function to be used for
// verifying authentication and authorization.
func RegisterSecurity(f RegisterSecurityFunc) func(next http.Handler) http.Handler {
	router, err := gorillamux.NewRouter(MustSwagger())
	if err != nil {
		panic("gorillamux.NewRouter: " + err.Error())
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route, pathParams, err := router.FindRoute(r)
			if err != nil {
				f.OnError(w, r, err)
				return
			}

			in := &openapi3filter.RequestValidationInput{
				Request:    r,
				PathParams: pathParams,
				Route:      route,
				Options: &openapi3filter.Options{
					AuthenticationFunc: f.Func,
				},
			}

			if err := validateSecurity(r.Context(), in); err != nil {
				f.OnError(w, r, err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func validateSecurity(ctx context.Context, input *openapi3filter.RequestValidationInput) error {
	security := input.Route.Operation.Security
	if security == nil {
		security = &input.Route.Spec.Security
		if security == nil {
			return nil
		}
	}

	return openapi3filter.ValidateSecurityRequirements(ctx, input, *security)
}
