package clarifai

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	clientID() string
	clientSecret() string
	accessToken() string
	apiHost() string
	apiPort() string
}

const VERSION = "V1"
const ROOT_URL = "api.clarifai.com"

type ClarifaiClient struct {
	clientID     string
	clientSecret string
	accessToken  string
}

// Initialize a new client object
func InitClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "unasigned"}
}

func (self *ClarifaiClient) post(values url.Values, endpoint string) ([]byte, error) {
	parts := []string{ROOT_URL, VERSION, endpoint}
	url := strings.Join(parts, "/")
	req, err := http.NewRequest("POST", url, string.NewReader(values.encode()))

	if err != nil {
		return nil, err
	}

	defer res.Body.close()

	body, err := ioutil.ReadAll(res.body)

	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 429 {
			return body, Error{"Throttled"}
		} else if res.StatusCode >= 400 && res.StatusCode < 500 {
			return body, Error{"Bad Request"}
		} else if res.StatusCode >= 500 && res.StatusCode < 600 {
			return body, Error{"Clarify Exception"}
		} else {
			return body, Error{"Unexpected Status Code"}
		}
	}

	return body, err
}

// clientID getter
func (self *ClarifaiClient) getClientID() string {
	return self.clientID
}

// clientSecret getter
func (self *ClarifaiClient) getClientSecret() string {
	return self.clientSecret
}

// accessToken getter
func (self *ClarifaiClient) getAccessToken() string {
	return self.accessToken
}
