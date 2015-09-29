package clarifai

import "strings"

// Configurations
const (
	Version         = "v1"
	RootURL         = "https://api.clarifai.com"
	TokenMaxRetries = 2
)

// Client contains scoped variables forindividual clients
type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Throttled    bool
	TokenRetries int
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *Client {
	return &Client{clientID, clientSecret, "", false, 0}
}

// SwitchThrottle is a convenience setter to switch the throttled flag
func (client *Client) SwitchThrottle() {
	client.Throttled = !client.Throttled
}

// Helper function to build URLs
func buildURL(endpoint string) string {
	parts := []string{RootURL, Version, endpoint}
	return strings.Join(parts, "/") + "/"
}
