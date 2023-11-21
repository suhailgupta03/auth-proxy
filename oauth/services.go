package oauth

// ServiceConfig holds OAuth configuration for a specific service.
type ServiceConfig struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	// Other fields like scopes, redirect URLs, etc.
}
