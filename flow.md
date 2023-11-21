```mermaid
sequenceDiagram
    participant C as Client
    participant E as Echo Server
    participant OCH as OAuth Callback Handler
    participant TS as Token Storage
    participant OP as OAuth Proxy
    participant SA as Slack API

    C->>E: Request /oauth/callback/:service with code
    E->>OCH: Handle callback
    OCH->>TS: Retrieve token for service
    alt Token exists
        TS-->>OCH: Return token
    else Token not found
        OCH->>OP: Exchange code for token
        OP->>SA: Request token
        SA-->>OP: Token response
        OP-->>OCH: Store new token
        OCH->>TS: Save token
    end
    OCH-->>E: Respond with token info
    E-->>C: OAuth callback response

    C->>E: Request /api/proxy/slack with URL and method
    E->>OP: Prepare API request
    OP->>TS: Get stored token for Slack
    TS-->>OP: Return token
    OP->>SA: Make API request with token
    SA-->>OP: API response
    OP-->>E: Forward API response
    E-->>C: API response to client

```

###

```mermaid
sequenceDiagram
    participant Client
    participant EchoServer as Echo Server
    participant OAuthHandler as OAuth Callback Handler
    participant TokenStorage
    participant OAuthProxy
    participant SlackAPI as Slack API

    Client->>+EchoServer: Request /oauth/callback/:service with code
    EchoServer->>+OAuthHandler: Handle callback
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
    OAuthHandler-->>-EchoServer: Respond with token info
    EchoServer-->>-Client: Display token or error message

    Client->>+EchoServer: Request /api/proxy/slack with URL
    EchoServer->>+OAuthProxy: Make API request
    OAuthProxy->>TokenStorage: Retrieve token for Slack
    TokenStorage-->>-OAuthProxy: Return token
    OAuthProxy->>+SlackAPI: GET or POST with Authorization header
    SlackAPI-->>-OAuthProxy: API response
    OAuthProxy-->>-EchoServer: Forward API response
    EchoServer-->>-Client: Display API response or error
```