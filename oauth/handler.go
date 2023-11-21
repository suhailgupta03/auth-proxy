package oauth

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/suhailgupta03/auth-proxy/tokenstorage"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func RegisterHandlers(e *echo.Echo, storage *tokenstorage.TokenStorage, proxy *OAuthProxy) {
	e.GET("/oauth/callback/:service", func(c echo.Context) error {
		service := c.Param("service")
		config, ok := Configs[service]
		if !ok {
			return c.String(http.StatusBadRequest, "Unknown service")
		}

		code := c.QueryParam("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "Missing code")
		}

		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to exchange token: %v", err))
		}

		// Store the token in the internal memory
		storage.Set(service, token)
		proxy.tokens[service] = token.AccessToken

		return c.JSON(http.StatusOK, map[string]interface{}{
			"access_token": token.AccessToken,
			"token_type":   token.TokenType,
			// Include any other token information you need in the response.
		})
	})
}

// In oauth/handler.go, add the following function

// RegisterAPIProxyHandler sets up a route to proxy API requests to Slack.
func RegisterAPIProxyHandler(e *echo.Echo, proxy *OAuthProxy) {
	// In oauth/handler.go, within the RegisterAPIProxyHandler function

	e.GET("/api/proxy/slack", func(c echo.Context) error {
		apiPath := c.QueryParam("url")
		method := c.QueryParam("method")
		if apiPath == "" {
			return c.String(http.StatusBadRequest, "Missing Slack API URL")
		}

		// Construct the full URL for the Slack API request
		fullURL := "https://slack.com/api/" + apiPath

		client := &http.Client{
			Transport: proxy,
		}
		req, err := http.NewRequest(http.MethodGet, fullURL, nil)
		if method == "POST" {
			postData := url.Values{}
			reqP, _ := http.NewRequest(http.MethodPost, fullURL, strings.NewReader(postData.Encode()))
			req = reqP
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)

		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to make proxied request: %v", err))
		}
		defer resp.Body.Close()

		// Forward the response body from Slack API to the client
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read response body: %v", err))
		}
		return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	})

}
