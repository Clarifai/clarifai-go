package clarifai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

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

// TokenResp is the expected response from /token/
type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "unasigned", false, 0, TokenMaxRetries}
}

func (client *ClarifaiClient) requestAccessToken() (*TokenResp, error) {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", client.clientID)
	form.Set("client_secret", client.clientSecret)
	formData := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", buildURL("token"), formData)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+client.accessToken)
	req.Header.Set("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	token := new(TokenResp)
	err = json.Unmarshal(body, token)
	fmt.Println(err)
	return token, err
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
