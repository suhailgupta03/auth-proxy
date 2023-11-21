package oauth

import (
	"errors"
	"net/http"
	"sync"
)

// OAuthProxy holds the configurations and tokens for different services.
type OAuthProxy struct {
	mu         sync.RWMutex
	configs    map[string]*ServiceConfig
	tokens     map[string]string // In-memory token storage, replace with secure storage.
	httpClient *http.Client
}

// NewOAuthProxy creates a new instance of OAuthProxy.
func NewOAuthProxy() *OAuthProxy {
	return &OAuthProxy{
		configs:    make(map[string]*ServiceConfig),
		tokens:     make(map[string]string),
		httpClient: &http.Client{},
	}
}

// RegisterService adds a new service configuration.
func (p *OAuthProxy) RegisterService(serviceName string, config *ServiceConfig) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.configs[serviceName] = config
}

// RoundTrip implements the http.RoundTripper interface.
// Note: The RoundTrip method implements the http.RoundTripper interface by adhering to the method
// signature defined by the interface and providing the necessary functionality to process an HTTP
// request and return an HTTP response.
func (p *OAuthProxy) RoundTrip(req *http.Request) (*http.Response, error) {
	serviceName := p.detectService(req)
	if serviceName == "" {
		return nil, errors.New("service detection failed")
	}

	p.mu.RLock()
	token, ok := p.tokens[serviceName]
	p.mu.RUnlock()

	if !ok {
		// Initiate OAuth2 flow to obtain token.
		// This is a placeholder for the actual OAuth2 implementation.
		token = "obtained_token"
		p.mu.Lock()
		p.tokens[serviceName] = token
		p.mu.Unlock()
	}

	print("Sending with token: ", token, req.URL.String(), req.Method)
	req.Header.Set("Authorization", "Bearer "+token)
	return p.httpClient.Do(req)
}

// detectService determines which service is being accessed.
func (p *OAuthProxy) detectService(req *http.Request) string {
	if req.URL.Host == "slack.com" {
		return "slack"
	}
	// Add more service detection logic as needed
	return ""
}

// refreshToken handles token refresh logic.
func (p *OAuthProxy) refreshToken(serviceName string) error {
	// Implement OAuth2 token refresh logic here.
	return nil
}
