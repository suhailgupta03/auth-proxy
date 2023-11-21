package tokenstorage

import (
	"golang.org/x/oauth2"
	"sync"
)

type TokenStorage struct {
	mu     sync.RWMutex
	tokens map[string]*oauth2.Token
}

func New() *TokenStorage {
	return &TokenStorage{
		tokens: make(map[string]*oauth2.Token),
	}
}

func (s *TokenStorage) Set(service string, token *oauth2.Token) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[service] = token
}

func (s *TokenStorage) Get(service string) (*oauth2.Token, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	token, exists := s.tokens[service]
	return token, exists
}
