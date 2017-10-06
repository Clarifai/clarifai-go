// Package clarifai provides a client interface to the Clarifai public API
package clarifai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	errUnexpectedStatusCode = errors.New("UNEXPECTED_STATUS_CODE")
	errClarifai             = errors.New("CLARIFAI_ERROR")
	errBadRequest           = errors.New("BAD_REQUEST")
	errThrottled            = errors.New("THROTTLED")
	errTokenInvalid         = errors.New("TOKEN_INVALID")
)

// Configurations
const (
	version = "v1"
	rootURL = "https://api.clarifai.com"
)

// Client contains scoped variables for individual clients
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
	token := new(TokenResp)
	if err := json.NewDecoder(res.Body).Decode(token); err != nil {
		return err
	}
	client.setAccessToken(token.AccessToken)
	return nil
}

func createRequestBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return json.Marshal(body)
}

func (client *Client) commonHTTPRequest(jsonBody interface{}, endpoint, verb string, retry bool, decodeInto interface{}) error {
	body, err := createRequestBody(jsonBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(verb, client.buildURL(endpoint), bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(body)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if client.Throttled {
			client.setThrottle(false)
		}
		if decodeInto != nil {
			if err := json.NewDecoder(res.Body).Decode(decodeInto); err != nil {
				return err
			}
		}
		return nil
	case http.StatusUnauthorized:
		if !retry {
			if err := client.requestAccessToken(); err != nil {
				return err
			}
			return client.commonHTTPRequest(jsonBody, endpoint, verb, true, decodeInto)
		}
		return errTokenInvalid
	case http.StatusTooManyRequests:
		client.setThrottle(true)
		return errThrottled
	case http.StatusBadRequest:
		return errBadRequest
	case http.StatusInternalServerError:
		return errClarifai
	default:
		return errUnexpectedStatusCode
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
