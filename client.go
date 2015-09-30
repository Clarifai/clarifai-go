package clarifai

import "strings"

// Configurations
const (
	Version                 = "v1"
	RootURL                 = "https://api.clarifai.com"
	MaxTokenRequestAttempts = 3
)

// Client contains scoped variables forindividual clients
type Client struct {
	APIHost              string
	ClientID             string
	ClientSecret         string
	AccessToken          string
	Throttled            bool
	TokenRequestAttempts int
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *Client {
	return &Client{RootURL, clientID, clientSecret, "", false, 0}
}

// SetThrottle is a convenience setter to switch the throttled flag
func (client *Client) SetThrottle(val bool) {
	client.Throttled = val
}

// Helper function to build URLs
func (client *Client) buildURL(endpoint string) string {
	parts := []string{client.APIHost, Version, endpoint}
	return strings.Join(parts, "/")
}
