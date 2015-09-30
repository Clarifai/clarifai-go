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

// TokenResp is the expected response from /token/
type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (client *Client) requestAccessToken() (*TokenResp, error) {
	if client.TokenRequestAttempts >= MaxTokenRequestAttempts {
		return nil, errors.New("MAX_TOKEN_REQUEST_ATTEMPTS_REACHED")
	}
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", client.ClientID)
	form.Set("client_secret", client.ClientSecret)
	formData := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", client.buildURL("token"), formData)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
	req.Header.Set("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		client.TokenRequestAttempts++
		return nil, errors.New("TOKEN_APP_INVALID")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	token := new(TokenResp)
	err = json.Unmarshal(body, token)
	if err != nil {
		return token, err
	}
	client.AccessToken = token.AccessToken
	return token, err
}
