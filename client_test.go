package clarifai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	ClientID     = "CLIENT_ID"
	ClientSecret = "CLIENT_SECRET"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)
}

func teardown() {
	server.Close()
}

func TestNewClarifaiClient(t *testing.T) {
	if client := NewClient(ClientID, ClientSecret); client == nil {
		t.Error("NewClient should not return nil")
	}
}

func TestNewClarifaiClientStoredValues(t *testing.T) {
	client := NewClient(ClientID, ClientSecret)
	if client.ClientID != ClientID || client.ClientSecret != ClientSecret {
		t.Error("NewClient should store the values of clientID and clientSecret")
	}
}

func TestRequestAccessToken(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"1234567890abcdefg","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`)
	})

	if err := client.requestAccessToken(); err != nil {
		t.Errorf("requestAccessToken() should not return an err upon success: %q", err)
	}
}

func TestRequestAccessTokenFail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	})

	if err := client.requestAccessToken(); err == nil {
		t.Errorf("requestAccessToken() should return an err with an invalid request: %v", err)
	}
}

func TestAccessTokenIsSaved(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"1234567890abcdefg","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`)
	})

	if err := client.requestAccessToken(); err != nil {
		t.Errorf("requestAccessToken() should not return err with a valid response")
	}

	if client.AccessToken != "1234567890abcdefg" {
		t.Errorf("requestAccessToken() should store the access token. Expected: 1234567890abcdefg, Got: %v", client.AccessToken)
	}
}
