package clarifai

import (
	"testing"
)

func TestClient(t *testing.T) {
	wantID := "testID"
	wantSecret := "testSecret"

	testClient := InitClient(wantID, wantSecret)

	gotID := testClient.getClientID()
	gotSecret := testClient.getClientSecret()

	if gotID != wantID {
		t.Errorf("InitClient(\"%q\") --> client.getClientID() == %q", wantID, gotID)
	}
	if gotSecret != wantSecret {
		t.Errorf("InitClient(\"%q\") --> client.getClientSecret() == %q", wantSecret, gotSecret)
	}
}
