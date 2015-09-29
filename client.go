package clarifai

import "strings"

// Configurations
const (
	Version         = "v1"
	RootURL         = "https://api.clarifai.com"
	TokenMaxRetries = 2
)

// ClarifaiClient contains scoped variables forindividual clients
type ClarifaiClient struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Throttled    bool
	TokenRetries int
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "", false, 0}
}

// Convenience setter to switch the throttled flag
func (client *ClarifaiClient) switchThrottle() {
	client.Throttled = !client.Throttled
}

// Helper function to build URLs
func buildURL(endpoint string) string {
	parts := []string{RootURL, Version, endpoint}
	return strings.Join(parts, "/") + "/"
}
