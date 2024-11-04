package credentials

import (
	"context"
)

// LocalTokenCache provides an interface to retrieve and save the access Token as a string.
// Interface allows integrators to fully control how access Token is caught on their side.
type LocalTokenCache interface {
	// RetrieveToken should return the access Token string.
	RetrieveToken(ctx context.Context) (string, error)

	// SaveToken should save the access Token string.
	SaveToken(ctx context.Context, token string) error
}
