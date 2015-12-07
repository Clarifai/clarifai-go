// Package clarifai provides a client interface to the Clarifai public API
package clarifai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Configurations
const (
	version = "v1"
	rootURL = "https://api.clarifai.com"
)

// Client contains scoped variables forindividual clients
type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	APIRoot      string
	Throttled    bool
}

// TokenResp is the expected response from /token/
type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// NewClient initializes a new Clarifai client
func NewClient(clientID, clientSecret string) *Client {
	return &Client{clientID, clientSecret, "unasigned", rootURL, false}
}

func (client *Client) requestAccessToken() error {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", client.ClientID)
	form.Set("client_secret", client.ClientSecret)
	formData := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", client.buildURL("token"), formData)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
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

	client.setAccessToken(token.AccessToken)
	return nil
}

func (client *Client) commonHTTPRequest(jsonBody interface{}, endpoint, verb string, retry bool) ([]byte, error) {
	if jsonBody == nil {
		jsonBody = struct{}{}
	}

	body, err := json.Marshal(jsonBody)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(verb, client.buildURL(endpoint), bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200, 201:
		if client.Throttled {
			client.setThrottle(false)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		return body, err
	case 401:
		if !retry {
			err := client.requestAccessToken()
			if err != nil {
				return nil, err
			}
			return client.commonHTTPRequest(jsonBody, endpoint, verb, true)
		}
		return nil, errors.New("TOKEN_INVALID")
	case 429:
		client.setThrottle(true)
		return nil, errors.New("THROTTLED")
	case 400:
		return nil, errors.New("ALL_ERROR")
	case 500:
		return nil, errors.New("CLARIFAI_ERROR")
	default:
		return nil, errors.New("UNEXPECTED_STATUS_CODE")
	}
}

// Helper function to build URLs
func (client *Client) buildURL(endpoint string) string {
	parts := []string{client.APIRoot, version, endpoint}
	return strings.Join(parts, "/")
}

// SetAccessToken will set accessToken to a new value
func (client *Client) setAccessToken(token string) {
	client.AccessToken = token
}

func (client *Client) setAPIRoot(root string) {
	client.APIRoot = root
}

func (client *Client) setThrottle(throttle bool) {
	client.Throttled = throttle
}
