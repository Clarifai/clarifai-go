package clarifai

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (client *Client) commonHTTPRequest(values url.Values, endpoint string) ([]byte, error) {
	req, err := http.NewRequest("POST", buildURL(endpoint), strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Length", strconv.Itoa(len(values.Encode())))
	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
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
		return nil, errors.New("TOKEN_INVALID")
	case 429:
		return nil, errors.New("THROTTLED: " + res.Header.Get("x-throttle-wait-seconds"))
	case 400:
		return nil, errors.New("ALL_ERROR")
	case 500:
		return nil, errors.New("CLARIFAI_ERROR")
	default:
		return nil, errors.New("UNEXPECTED_STATUS_CODE")
	}
}
