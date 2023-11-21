package main

import (
	"github.com/labstack/echo/v4"
	"github.com/suhailgupta03/auth-proxy/oauth"
	"github.com/suhailgupta03/auth-proxy/tokenstorage"
	"golang.org/x/oauth2"
)

var (
	// Global variable to hold OAuth configurations for different services.
	oauthConfigs = map[string]*oauth2.Config{
		// Add OAuth configurations for each service.
		"slack": {
			ClientID:     "2314607060.5796126504288",
			ClientSecret: "00ed067137f8a5a70ee19eca14df7960",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://slack.com/oauth/v2/authorize",
				TokenURL: "https://slack.com/api/oauth.v2.access",
			},
			RedirectURL: "https://4784-116-212-183-134.ngrok-free.app/oauth/callback/slack",
			Scopes:      []string{"channels:read", "chat:write", "commands", "groups:history", "im:write", "users:read"},
		},
		// Add more services as needed.
	}
)

func main() {
	e := echo.New()

	// Initialize token storage
	tokenStorage := tokenstorage.New()

	oauthProxy := oauth.NewOAuthProxy()
	oauth.RegisterService(oauthProxy, "slack")

	// Register OAuth callback handler
	oauth.RegisterHandlers(e, tokenStorage, oauthProxy)
	oauth.RegisterAPIProxyHandler(e, oauthProxy)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
	// Here's the typical sequence of events:
	// An http.Client is created with its Transport field set to an instance of OAuthProxy.
	// A new HTTP request is constructed and client.Do(request) is called.
	// The Do method internally calls the RoundTrip method of the Transport (which is our OAuthProxy).
	// The RoundTrip method processes the request, adding authentication headers or performing other actions as necessary.
	// The RoundTrip method then delegates the actual network round trip to the default transport or a custom transport if one is provided within the OAuthProxy.
	// Once the response is received, the RoundTrip method returns it back up the call stack to where client.Do(request) was called.

}
