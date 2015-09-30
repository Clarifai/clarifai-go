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
		fmt.Fprintln(w, `{"access_token":"U5j3rECLLZ86pWRjK35g489QJ4zrQI","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`)
	})

	_, err := client.requestAccessToken()

	if err != nil {
		t.Errorf("requestAccessToken() should not return an err upon success: %v", err)
	}
}
