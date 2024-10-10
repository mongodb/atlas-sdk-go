package httpdigest

import (
	"net/http"

	"github.com/mongodb-forks/digest"
)

// TODO digest auth module
// CreateDigestClient
func CreateDigestClient(apiKey, apiSecret string)  (*http.Client, error) {
	transport := digest.NewTransport(apiKey, apiSecret)
	httpClient, err := transport.Client()
	if err != nil {
		return nil, err;
	}
	return httpClient, nil;
}

