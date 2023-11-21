package oauth

import (
	"fmt"
	"golang.org/x/oauth2"
)

var Configs = map[string]*oauth2.Config{
	// Add OAuth configurations for each service.
	"slack": {
		ClientID:     "2314607060.5796126504288",
		ClientSecret: "00ed067137f8a5a70ee19eca14df7960",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://slack.com/oauth/v2/authorize",
			TokenURL: "https://slack.com/api/oauth.v2.access",
		},
		RedirectURL: "https://4784-116-212-183-134.ngrok-free.app/oauth/callback/slack",
		Scopes:      []string{"channels:read"},
	},
	// Add more services as needed.
}

func RegisterService(proxy *OAuthProxy, serviceName string) error {
	config, exists := Configs[serviceName]
	if !exists {
		return fmt.Errorf("OAuth configuration for service %s does not exist", serviceName)
	}
	proxy.RegisterService(serviceName, &ServiceConfig{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		AuthURL:      config.Endpoint.AuthURL,
		TokenURL:     config.Endpoint.TokenURL,
		// Add other fields as necessary from the config.
	})
	return nil
}
