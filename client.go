package clarifai

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client interface {
	getClientID() string
	getClientSecret() string
	requestAccessToken() string
}

const VERSION = "v1"
const ROOT_URL = "https://api.clarifai.com"

const TOKEN_MAX_RETRIES = 2

type ClarifaiClient struct {
	clientID        string
	clientSecret    string
	accessToken     string
	throttled       bool
	tokenRetries    int
	tokenMaxRetries int
}

type TokenResp struct {
	access_token string `json:"access_token"`
	expires_in   int    `json:"expires_in"`
	scope        string `json:"scope"`
	token_type   string `json:"token_type"`
}

// Initialize a new client object
func NewClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "unasigned", false, 0, TOKEN_MAX_RETRIES}
}

func (client *ClarifaiClient) requestAccessToken() (*TokenResp, error) {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", client.clientID)
	form.Set("client_secret", client.clientSecret)
	formData := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", buildUrl("token"), formData)

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
func buildUrl(endpoint string) string {
	parts := []string{ROOT_URL, VERSION, endpoint}
	return strings.Join(parts, "/") + "/"
}
