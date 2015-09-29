package clarifai

import "testing"

const (
	ClientID     = "CLIENT_ID"
	ClientSecret = "CLIENT_SECRET"
)

func TestNewClarifaiClient(t *testing.T) {
	client := NewClient(ClientID, ClientSecret)
	if client == nil {
		t.Error("NewClient should not return nil")
	}
}

func TestNewClarifaiClientStoredValues(t *testing.T) {
	client := NewClient(ClientID, ClientSecret)
	if client.ClientID != ClientID || client.ClientSecret != ClientSecret {
		t.Error("NewClient should store the values of clientID and clientSecret")
	}
}
