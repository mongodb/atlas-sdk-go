package credentials

import (
	"context"
	"errors"
	"sync"
)

// LocalTokenCache provides an interface to retrieve and save the access Token as a string.
// Interface allows integrators to fully control how access Token is catched on their side.
type LocalTokenCache interface {
	// RetrieveToken should return the access Token string.
	RetrieveToken(ctx context.Context) (*string, error)

	// SaveToken should save the access Token string.
	SaveToken(ctx context.Context, token string) error
}

// InMemoryTokenCache is an default implementation of LocalTokenCache that stores the Token in memory.
// Implementation provides locking mechanism in order to prevent overriding access tokens
type InMemoryTokenCache struct {
	token *string
	mu    sync.Mutex
}

func (s *InMemoryTokenCache) RetrieveToken(_ context.Context) (*string, error) {
	// Locking will avoid Token overrides when sharing TokenCache in multiple threads
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.token != nil {
		return s.token, nil
	}
	return nil, errors.New("InMemoryTokenCache: Token not found. Token needs to be refreshed in backend")
}

func (s *InMemoryTokenCache) SaveToken(_ context.Context, token string) error {
	// Locking will avoid Token overrides when sharing TokenCache in multiple threads
	s.mu.Lock()
	defer s.mu.Unlock()
	s.token = &token
	return nil
}
