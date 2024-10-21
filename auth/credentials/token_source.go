package credentials

import (
	"errors"
	"sync"
)

// TokenSource provides an interface to retrieve and save the access token as a string.
// Interface allows integrators to fully control how access token is catched on their side.
type TokenSource interface {
	// RetrieveToken should return the access token string.
	RetrieveToken() (*string, error)

	// SaveToken should save the access token string.
	SaveToken(token string) error
}

// InMemoryTokenSource is an default implementation of TokenSource that stores the token in memory.
type InMemoryTokenSource struct {
	token *string
	mu    sync.Mutex
}

func (s *InMemoryTokenSource) RetrieveToken() (*string, error) {
	// Locking will avoid token overrides when sharing client in multiple threads
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.token != nil {
		return s.token, nil
	}
	return nil, errors.New("InMemoryTokenSource: token not found. Token needs to be refreshed in backend")
}

func (s *InMemoryTokenSource) SaveToken(token string) error {
	// Locking will avoid token overrides when sharing client in multiple threads
	s.mu.Lock()
	defer s.mu.Unlock()
	s.token = &token
	return nil
}
