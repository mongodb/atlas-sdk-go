package admin // import "go.mongodb.org/atlas-sdk/v20240805005/admin"

import (
	"errors"
	"net/http"
	"strings"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas-sdk/v20240805005/internal/core"
)

const (
	// Version the version of the current API client inherited from.
	Version = core.Version
)

// NewClient returns a new API Client.
func NewClient(modifiers ...ClientModifier) (*APIClient, error) {
	defaultConfig := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: core.DefaultCloudURL,
		},
		},
		UserAgent: core.DefaultUserAgent,
	}
	for _, modifierFunction := range modifiers {
		err := modifierFunction(defaultConfig)
		if err != nil {
			return nil, err
		}
	}

	return NewAPIClient(defaultConfig), nil
}

// ClientModifiers lets you create function that controls configuration before creating client.
type ClientModifier func(*Configuration) error

// UseDigestAuth provides Digest authentication for Go SDK.
// UseDigestAuth is provided as helper to create a default HTTP client that supports HTTP Digest authentication.
// Warning: any previously set httpClient will be overwritten. To fully customize HttpClient use UseHTTPClient method.
func UseDigestAuth(apiKey, apiSecret string) ClientModifier {
	return func(c *Configuration) error {
		transport := digest.NewTransport(apiKey, apiSecret)
		httpClient, err := transport.Client()
		if err != nil {
			return err
		}
		c.HTTPClient = httpClient
		return nil
	}
}

// Advanced modifiers.

// UseHTTPClient set custom http client implementation.
//
// Warning: UseHTTPClient overrides any previously set httpClient including the one set by UseDigestAuth.
// To set a custom http client with HTTP diggest support use:
//
//	transport := digest.NewTransportWithHTTPRoundTripper(apiKey, apiSecret, yourHttpTransport)
//	client := UseHTTPClient(transport.Client())
func UseHTTPClient(client *http.Client) ClientModifier {
	return func(c *Configuration) error {
		c.HTTPClient = client
		return nil
	}
}

// UseDebug enable debug level logging.
func UseDebug(debug bool) ClientModifier {
	return func(c *Configuration) error {
		c.Debug = debug
		return nil
	}
}

// UseBaseURL set custom base url. If empty, default is used.
func UseBaseURL(baseURL string) ClientModifier {
	return func(c *Configuration) error {
		if baseURL == "" {
			baseURL = core.DefaultCloudURL
		}
		urlWithoutSuffix := strings.TrimSuffix(baseURL, "/")
		c.Servers = ServerConfigurations{ServerConfiguration{
			URL: urlWithoutSuffix,
		}}
		return nil
	}
}

// UseUserAgent set custom UserAgent header.
func UseUserAgent(userAgent string) ClientModifier {
	return func(c *Configuration) error {
		c.UserAgent = userAgent
		return nil
	}
}

// AsError checks if API returned known error type.
func AsError(err error) (*ApiError, bool) {
	var openapiError *GenericOpenAPIError
	if ok := errors.As(err, &openapiError); !ok {
		return nil, false
	}
	errModel := openapiError.Model()
	return &errModel, true
}

// IsErrorCode returns true if the error contains the specific code.
func IsErrorCode(err error, code string) bool {
	mappedErr, _ := AsError(err)
	if mappedErr == nil {
		return false
	}
	return mappedErr.GetErrorCode() == code
}
