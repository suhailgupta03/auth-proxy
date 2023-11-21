package main

import (
	"golang.org/x/oauth2"
	"sync"
)

// TokenStorage is a thread-safe map to store access tokens in memory.
type TokenStorage struct {
	mu     sync.RWMutex
	tokens map[string]*oauth2.Token
}

// NewTokenStorage initializes a new TokenStorage instance.
func NewTokenStorage() *TokenStorage {
	return &TokenStorage{
		tokens: make(map[string]*oauth2.Token),
	}
}

// Set stores a token associated with a service.
func (s *TokenStorage) Set(service string, token *oauth2.Token) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[service] = token
}

// Get retrieves a token associated with a service.
func (s *TokenStorage) Get(service string) (*oauth2.Token, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	token, exists := s.tokens[service]
	return token, exists
}
