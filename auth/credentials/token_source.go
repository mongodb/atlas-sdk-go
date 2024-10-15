package credentials

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// TokenSource provides an interface to retrieve and save the access token as a string.
// Interface allows integrators to fully control how access token is catched on their side.
type TokenSource interface {
    // RetrieveToken should return the access token string.
    RetrieveToken() (string, error)

    // SaveToken should save the access token string.
    SaveToken(token string) error
}

// InMemoryTokenSource is an implementation of TokenSource that stores the token in memory.
type InMemoryTokenSource struct {
	tkn *token
	mu  sync.Mutex
}

func (s *InMemoryTokenSource) RetrieveToken() (*token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tkn != nil {
		return s.tkn, nil
	}
	return nil, errors.New("token not found")
}

func (s *InMemoryTokenSource) SaveToken(tkn *token) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tkn = tkn
	return nil
}

// TODO remove/move to example

// FileTokenSource is an implementation of TokenSource that stores the token in a file.
type FileTokenSource struct {
	filePath string
	mu       sync.Mutex
}

func (s *FileTokenSource) RetrieveToken() (*token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileData, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var tkn token
	err = json.Unmarshal(fileData, &tkn)
	if err != nil {
		return nil, err
	}

	return &tkn, nil
}

func (s *FileTokenSource) SaveToken(tkn *token) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileData, err := json.Marshal(tkn)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, fileData, 0600)
}
