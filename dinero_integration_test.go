// +build integration

package dinero

import (
	"os"
	"testing"
)

func TestClient_Authorize(t *testing.T) {
	key := os.Getenv("CLIENTKEY")
	secret := os.Getenv("CLIENTSECRET")
	apiKey := os.Getenv("CLIENTAPIKEY")

	c := NewClient(key, secret)

	err := c.Authorize(apiKey)
	if err != nil {
		t.Errorf("Error occured when trying to talk to dinero auth api: %v", err)
	}

	var defaultString string
	if c.token == defaultString {
		t.Errorf("The client did not set the token correctly")
	}
}
