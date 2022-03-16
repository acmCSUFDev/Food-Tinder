// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version (devel) DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/render"
)

// Error object returned on any error
type Error struct {
	Message string `json:"message"`
}

// Snowflake ID
type ID int64

// Optional metadata included on login
type LoginMetadata struct {
	// The User-Agent used for logging in
	UserAgent *string `json:"user_agent,omitempty"`
}

// Session defines model for Session.
type Session struct {
	Expiry time.Time `json:"expiry"`

	// Optional metadata included on login
	Metadata *LoginMetadata `json:"metadata,omitempty"`
	Token    string         `json:"token"`

	// Snowflake ID
	UserID ID `json:"user_id"`
}

// LoginParams defines parameters for Login.
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
type Response struct {
	body        interface{}
	statusCode  int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.statusCode)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(statusCode int) *Response {
	resp.statusCode = statusCode
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// LoginJSON200Response is a constructor method for a Login response.
// A *Response is returned with the configured status code and content type from the spec.
func LoginJSON200Response(body Session) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// LoginJSON401Response is a constructor method for a Login response.
// A *Response is returned with the configured status code and content type from the spec.
func LoginJSON401Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  401,
		contentType: "application/json",
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUQW/bPAz9KwK/76gk7lbs4FuBbkCADhvQ7lQUg2YxrlqbUkW5rRH4vw+U4zaNg/ay",
	"UwiFfOR7j/QWKt8GT0iJodwCV7fYmhx+jdFHCSxyFV1IzhOU47Pyf+6wSipi6iKhVZ6UoV5hrtEQog8Y",
	"k8OM1CKzqVHC1AeEEjhFRzUMg4aID52LaKG8fkm8GTSsz+e9L8k/bRpzj2p9DvoATMPzovaL3aOj9OUU",
	"Bg0Xvnb0HZOxJpk55I8cmEa1uxTlqGo6O3JqpHjGp2OMv02NlOZ4V7eofjHGxZn8rzpGqzY+ClLtqFYZ",
	"7lCFQcMlMmeE7UEzfA4u9hJtfGxNghKsSbhIrsU5lIZ2j+r/ETdQwn+rV5NXO4dXb3UZNCR/j3TEIz3y",
	"dfYjxPX5zNCpckLXE52bTNrRxueOLjWYzUA6+7lWI2LW7Zv3dnHlyKLs1SPGUSUolsXyRIbzAckEByV8",
	"XhbLQswy6TZLtxrtE0k9Z69EWCNOrS2U42rkgmhaTBgZyustSAk8dBh70ECmlcmESA73+aXYod6dzFHh",
	"OPWZlzgHgz4OHQzzk4/2XegX7/eyP2h3I3gcPPG4SJ+KQn4qT2m3uSaExlVZj9Udj8v32vI9q6dtzS6+",
	"PYALX9dolSPFXVUh86Zrml7GOy1O/tkA48fpSPs1PZrGWTUZtnoRTHK5a1sj1yRjyowdy1FOycqQVXsF",
	"w/A3AAD//wT14JogBQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}