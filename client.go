package clarifai

import "strings"

// Client is the interface to publicly exported functions
type Client interface {
	getClientID() string
	getClientSecret() string
	requestAccessToken() string
}

// Configurations
const (
	Version         = "v1"
	RootURL         = "https://api.clarifai.com"
	TokenMaxRetries = 2
)

// ClarifaiClient contains scoped variables forindividual clients
type ClarifaiClient struct {
	clientID        string
	clientSecret    string
	accessToken     string
	throttled       bool
	tokenRetries    int
	tokenMaxRetries int
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "unasigned", false, 0, TokenMaxRetries}
}

// clientID getter
func (client *ClarifaiClient) getClientID() string {
	return client.clientID
}

// clientSecret getter
func (client *ClarifaiClient) getClientSecret() string {
	return client.clientSecret
}

// accessToken getter
func (client *ClarifaiClient) getAccessToken() string {
	return client.accessToken
}

// Determine if the client is currently being throttled by the host
func (client *ClarifaiClient) isThrottled() bool {
	return client.throttled
}

// Convenience setter to switch the throttled flag
func (client *ClarifaiClient) switchThrottle() {
	client.throttled = !client.throttled
}

// Helper function to build URLs
func buildURL(endpoint string) string {
	parts := []string{RootURL, Version, endpoint}
	return strings.Join(parts, "/") + "/"
}
