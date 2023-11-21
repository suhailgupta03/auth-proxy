```mermaid
sequenceDiagram
    participant Client
    participant AuthService as "Auth Service"
    box "Auth Service" #LightBlue
        participant EchoServer as Echo Server
        participant OAuthHandler as OAuth Callback Handler
        participant TokenStorage
        participant OAuthProxy
    end
    participant SlackAPI as Slack API

    Client->>+AuthService: Request /oauth/callback/:service with code
    AuthService->>+OAuthHandler: Handle callback
    OAuthHandler->>+TokenStorage: Get stored token for service
    alt Token found
        TokenStorage-->>-OAuthHandler: Return token
    else Token not found
        OAuthHandler->>+OAuthProxy: Exchange code for token
        OAuthProxy->>+SlackAPI: POST to TokenURL
        SlackAPI-->>-OAuthProxy: Access token
        OAuthProxy-->>-OAuthHandler: Store and return token
        OAuthHandler->>TokenStorage: Save token
    end
    OAuthHandler-->>-AuthService: Respond with token info
    AuthService-->>-Client: Display token or error message

    Client->>+AuthService: Request /api/proxy/slack with URL
    AuthService->>+OAuthProxy: Make API request
    OAuthProxy->>TokenStorage: Retrieve token for Slack
    TokenStorage-->>-OAuthProxy: Return token
    OAuthProxy->>+SlackAPI: GET or POST with Authorization header
    SlackAPI-->>-OAuthProxy: API response
    OAuthProxy-->>-AuthService: Forward API response
    AuthService-->>-Client: Display API response or error


```

###
