package clarifai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	clientID     = "client_id"
	clientSecret = "client_secret"
)

func TestRequestAccessToken(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(clientID, clientSecret)
	client.APIHost = server.URL

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"1234567890abcdefg","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`)
	})

	_, err := client.requestAccessToken()

	if err != nil {
		t.Errorf("requestAccessToken() should not return an err upon success: %v", err)
	}
}

func TestRequestAccessTokenFail(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(clientID, clientSecret)
	client.APIHost = server.URL

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	})

	token, _ := client.requestAccessToken()

	if token != nil {
		t.Errorf("requestAccessToken() should return an err with an invalid request: %v", token)
	}
}

func TestAccessTokenIsSaved(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(clientID, clientSecret)
	client.APIHost = server.URL

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"1234567890abcdefg","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`)
	})

	_, err := client.requestAccessToken()

	if err != nil {
		if client.AccessToken != "1234567890abcdefg" {
			t.Errorf("requestAccessToken() should store the access token. Expected: 1234567890abcdefg, Got: %v", client.AccessToken)
		}
	}
}
