package clarifai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInfo(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)

	defer server.Close()

	mux.HandleFunc("/v1/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"status_code":"OK","status_msg":"All images in request have completed successfully. ","results":{"max_image_size":100000,"default_language":"en","max_video_size":100000,"max_image_bytes":10485760,"min_image_size":1,"default_model":"default","max_video_bytes":104857600,"max_video_duration":1800,"max_batch_size":128,"max_video_batch_size":1,"min_video_size":1,"api_version":0.1}}`)
	})

	_, err := ClarifaiInfo(client)

	if err != nil {
		t.Errorf("requestAccessToken() should not return an err upon success: %v", err)
	}
}
