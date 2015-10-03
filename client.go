package clarifai

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client is the interface to exported functions
type Client interface {
	ClientID() string
	ClientSecret() string
	AccessToken() string
	Throttled() bool
	APIRoot() string
	buildURL(string) string
	setAPIRoot(string)
	commonHTTPRequest(url.Values, string, string, bool) ([]byte, error)
}

// Configurations
const (
	Version         = "v1"
	RootURL         = "https://api.clarifai.com"
	TokenMaxRetries = 2
)

// ClarifaiClient contains scoped variables forindividual clients
type ClarifaiClient struct {
	clientID     string
	clientSecret string
	accessToken  string
	apiRoot      string
	throttled    bool
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
	return &ClarifaiClient{clientID, clientSecret, "unasigned", RootURL, false}
}

func (client *ClarifaiClient) requestAccessToken() error {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", client.ClientID())
	form.Set("client_secret", client.ClientSecret())
	formData := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", client.buildURL("token"), formData)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+client.AccessToken())
	req.Header.Set("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	token := new(TokenResp)
	err = json.Unmarshal(body, token)

	if err != nil {
		return err
	}

	client.SetAccessToken(token.AccessToken)
	return nil
}

func (client *ClarifaiClient) commonHTTPRequest(values url.Values, endpoint, verb string, retry bool) ([]byte, error) {
	if values == nil {
		values = url.Values{}
	}

	req, err := http.NewRequest(verb, client.buildURL(endpoint), strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Length", strconv.Itoa(len(values.Encode())))
	req.Header.Set("Authorization", "Bearer "+client.AccessToken())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200, 201:
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		return body, err
	case 401:
		if !retry {
			err := client.requestAccessToken()
			if err != nil {
				return nil, err
			}
			return client.commonHTTPRequest(values, endpoint, verb, true)
		}
		return nil, errors.New("TOKEN_INVALID")
	case 429:
		client.setThrottle(true)
		return nil, errors.New("THROTTLED: " + res.Header.Get("x-throttle-wait-seconds"))
	case 400:
		return nil, errors.New("ALL_ERROR")
	case 500:
		return nil, errors.New("CLARIFAI_ERROR")
	default:
		return nil, errors.New("UNEXPECTED_STATUS_CODE")
	}
}

// Helper function to build URLs
func (client *ClarifaiClient) buildURL(endpoint string) string {
	parts := []string{client.APIRoot(), Version, endpoint}
	return strings.Join(parts, "/")
}

// ClientID will return the clientID
func (client *ClarifaiClient) ClientID() string {
	return client.clientID
}

// ClientSecret will return the clientSecret
func (client *ClarifaiClient) ClientSecret() string {
	return client.clientSecret
}

// AccessToken will return the current accessToken
func (client *ClarifaiClient) AccessToken() string {
	return client.accessToken
}

// SetAccessToken will set accessToken to a new value
func (client *ClarifaiClient) SetAccessToken(token string) {
	client.accessToken = token
}

// Throttled returns the state of the client being throttled or not
func (client *ClarifaiClient) Throttled() bool {
	return client.throttled
}

// APIRoot returns the root path for the API
func (client *ClarifaiClient) APIRoot() string {
	return client.apiRoot
}

func (client *ClarifaiClient) setAPIRoot(root string) {
	client.apiRoot = root
}

func (client *ClarifaiClient) setThrottle(throttle bool) {
	client.throttled = throttle
}
