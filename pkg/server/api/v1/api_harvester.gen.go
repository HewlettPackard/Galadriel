// Package api_harvester provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api_harvester

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for FederationRelationshipStatus.
const (
	FederationRelationshipStatusActive   FederationRelationshipStatus = "active"
	FederationRelationshipStatusInactive FederationRelationshipStatus = "inactive"
	FederationRelationshipStatusInvited  FederationRelationshipStatus = "invited"
)

// Defines values for TrustBundleStatus.
const (
	TrustBundleStatusActive   TrustBundleStatus = "active"
	TrustBundleStatusInactive TrustBundleStatus = "inactive"
	TrustBundleStatusToDelete TrustBundleStatus = "to_delete"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// FederationRelationship defines model for FederationRelationship.
type FederationRelationship struct {
	FederationGroupId               int64                         `json:"federationGroupId"`
	Id                              int64                         `json:"id"`
	SpireServer                     string                        `json:"spireServer"`
	SpireServerConsent              *string                       `json:"spireServerConsent,omitempty"`
	SpireServerFederatedWith        string                        `json:"spireServerFederatedWith"`
	SpireServerFederatedWithConsent *string                       `json:"spireServerFederatedWithConsent,omitempty"`
	Status                          *FederationRelationshipStatus `json:"status,omitempty"`
}

// FederationRelationshipStatus defines model for FederationRelationship.Status.
type FederationRelationshipStatus string

// TrustBundle defines model for TrustBundle.
type TrustBundle struct {
	Bundle      string             `json:"bundle"`
	Id          int64              `json:"id"`
	Status      *TrustBundleStatus `json:"status,omitempty"`
	TrustDomain *string            `json:"trustDomain,omitempty"`
}

// TrustBundleStatus defines model for TrustBundle.Status.
type TrustBundleStatus string

// GetFederationRelationshipsParams defines parameters for GetFederationRelationships.
type GetFederationRelationshipsParams struct {
	// filter relationships by spireServer
	SpireServer *string `form:"spireServer,omitempty" json:"spireServer,omitempty"`

	// filter relationships by status
	Status *string `form:"status,omitempty" json:"status,omitempty"`

	// filter relationships by status
	FederationGroupId *int64 `form:"federationGroupId,omitempty" json:"federationGroupId,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /FederationRelationship)
	GetFederationRelationships(ctx echo.Context, params GetFederationRelationshipsParams) error

	// (GET /FederationRelationship/{relationshipID})
	GetRelationshipbyID(ctx echo.Context, relationshipID int64) error

	// (PUT /FederationRelationship/{relationshipID})
	UpdateFederatedRelationshipStatus(ctx echo.Context, relationshipID int64) error

	// (PUT /trustBundles/{trustBundleId})
	UpdateTrustBundle(ctx echo.Context, trustBundleId int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFederationRelationships converts echo context to params.
func (w *ServerInterfaceWrapper) GetFederationRelationships(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFederationRelationshipsParams
	// ------------- Optional query parameter "spireServer" -------------

	err = runtime.BindQueryParameter("form", true, false, "spireServer", ctx.QueryParams(), &params.SpireServer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter spireServer: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Optional query parameter "federationGroupId" -------------

	err = runtime.BindQueryParameter("form", true, false, "federationGroupId", ctx.QueryParams(), &params.FederationGroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter federationGroupId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFederationRelationships(ctx, params)
	return err
}

// GetRelationshipbyID converts echo context to params.
func (w *ServerInterfaceWrapper) GetRelationshipbyID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "relationshipID" -------------
	var relationshipID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "relationshipID", runtime.ParamLocationPath, ctx.Param("relationshipID"), &relationshipID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter relationshipID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRelationshipbyID(ctx, relationshipID)
	return err
}

// UpdateFederatedRelationshipStatus converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateFederatedRelationshipStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "relationshipID" -------------
	var relationshipID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "relationshipID", runtime.ParamLocationPath, ctx.Param("relationshipID"), &relationshipID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter relationshipID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateFederatedRelationshipStatus(ctx, relationshipID)
	return err
}

// UpdateTrustBundle converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTrustBundle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "trustBundleId" -------------
	var trustBundleId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "trustBundleId", runtime.ParamLocationPath, ctx.Param("trustBundleId"), &trustBundleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter trustBundleId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTrustBundle(ctx, trustBundleId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/FederationRelationship", wrapper.GetFederationRelationships)
	router.GET(baseURL+"/FederationRelationship/:relationshipID", wrapper.GetRelationshipbyID)
	router.PUT(baseURL+"/FederationRelationship/:relationshipID", wrapper.UpdateFederatedRelationshipStatus)
	router.PUT(baseURL+"/trustBundles/:trustBundleId", wrapper.UpdateTrustBundle)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RXwW7jNhD9FWJaoBdl5SaLotBt02xT34Kkix4WQUGLI5kLiWSGI2fdQP9ekLJsKZYT",
	"Z4sN0lMkeYZ88+bNI/MAua2dNWjYQ/YAPl9iLePjRyJL4cGRdUisMX7OrcLwt7BUS4YMtOGzU0iA1w67",
	"VyyRoE2gRu9lGaM3P3ombUpo2wQI7xpNqCD73K25i7/dLmYXXzDnsNbvqJAka2uusYp//VK7fXjFNu6S",
	"bOPm6jHWX95PYtXHBnqnCW+QVkijjE1pyeNSRxm/WePR8MsTN+Wj+kvz8j+mvxAES24itWiaOnRL5qxX",
	"oV/aDB5XmlENWneg1VpBMtGkMa9PlD6ljT+p8XzeGFXhviAW2+/bYhdrRj9V6/EiOIoUtn8rrJBxgpYE",
	"OKC+sLXU5qhWTBEZ11DdGklf6j5FIVebwoZ9FPqctAvkQwYfruaCrVigaDwqUVgSoUwKNZhS3Gteij8k",
	"rdAzkn8XttQc6ISbq/n1R3FOWpUoTsSlrKQijVUXTiFefLiaQwLhrdvt53ezd7NQu3VopNOQwVn8lICT",
	"vIyEpocnvUTer+AauSHjhawqUfRSETRI9cIWQooOsI+aEtoIXuIAte+VF7QTU4N1wCXyNB4fMZOsMRAD",
	"2efHuApdBQrGQBZrMdZ56D3cNUhrSMDIOjZ8PAnRkY9UyNEgOv0e2L//8fW3njSGfRSHBrO9DTPiXTC4",
	"KKbT2aw7swxvHE86V+k8bpF+8dbsDr3wpBnrmPgjYQEZ/JDujsd0czYeEmi7xSOJ5LqbujEnJfKYkJ+8",
	"6PFCjC5kU/GLID+FtDvCJ4A0Br86zMOkYB/TJodmL30Ygp5ftAeHMRSoJMtoJNagsFRKo/+JyVPDNdxm",
	"sZ5fPDdVcxWGOYzuENLGwgiZNK5Q9eoKprIT17gIGNopU4Ovq7RvEdjzgnpLekrANRMSaZySjEJ2Bty5",
	"QWfQo44u1iL2aKyYTzF5eyMY0nPT+8o36qfD9VrKuWvQ87lV6yd68/Xk/v7+JCx40lCFJtyS1fdQ0WZ/",
	"31NjqewvBZEUBe2e2t/vd9ZY0VfyFryMd7dCnz4M3uYqGtikOj+5ykolpBjeKadVOI44UnURhuguas+p",
	"bgT5fyK6ISlHKO0AHe13NNdnEAZHHaJ6Uyd0+M8j3gs7kTVUQQZLZpelaWVzWS2t5+zs9Gz2awrtbftv",
	"AAAA//8vq4HA3g8AAA==",
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
