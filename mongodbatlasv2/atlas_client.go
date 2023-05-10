package mongodbatlasv2 // import "go.mongodb.org/atlas/mongodbatlas/v2"

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas-sdk/core"
)

const (
	// CloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com"
	// Version the version of the current API client inherited from.
	Version = core.Version
	// Name of the v2 API client.
	ClientName = "go-atlas-sdk-admin"
)

// NewClient returns a new MongoDBAtlas API Client.
func NewClient(modifiers ...ClientModifier) (*APIClient, error) {
	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	defaultConfig := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: DefaultCloudURL,
		},
		},
		UserAgent: userAgent,
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

// UseHttpClient set custom http client implementation.
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

// UseBaseURL set custom base url.
func UseBaseURL(baseURL string) ClientModifier {
	return func(c *Configuration) error {
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
func AsError(err error) (*Error, bool) {
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
