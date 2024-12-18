// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

// HealthResponse defines model for HealthResponse.
type HealthResponse struct {
	Status    *string    `json:"status,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Health Check
	// (GET /health)
	GetHealth(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// GetHealth operation middleware
func (siw *ServerInterfaceWrapper) GetHealth(c *fiber.Ctx) error {

	return siw.Handler.GetHealth(c)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Get(options.BaseURL+"/health", wrapper.GetHealth)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5RSTW/UMBD9K9bAMWwCFRffEAdacSgqx4rD4EwSt/HYsicVq1X+OxonbcVWourJo+f5",
	"eG/mncDFkCITSwF7guImCljDS8JZphsqKXIhRVKOibJ42lIFZakR/cGQZgIL19+hATkmjYtkzyOsDYgP",
	"VARD0uQh5oACFnoU+qBfL0vWJyT+viMnsCrkeYjaoafisk/iI4PdaRo3kbs3xH2KnqWYIWYjExlcZCIW",
	"71DzTaH84F0d6aVS3uu/1vovP66ggQfKZWv+8dAdOpUQEzEmDxYuKtRAXCQtqsOzUGac25GYMgr17VR7",
	"tpj8YSQ+jBEaSOjucdSJ22+FZKoL3As0HElearwhWTKXKmhLNdv2TRwq+CxLL1SlXvVg4RvJ5eO0vF+y",
	"TvzUdfq4yEJcJ2JK876l9q7o2Ec3aPQ+0wAW3rXPdml3r7RnRqm3+pf/z42e8WWnf9Sdfu4u3sQB5/l6",
	"AHv7NjbNuW8DlVIPcXrddr/+L2bhJzlrA2UJAfPxzFNqNRwL2NsdBm26/g0AAP//PMx1PHkDAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
