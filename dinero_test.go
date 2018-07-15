package dinero

import (
	"os"
	"strconv"
	"testing"
)

func TestClient_Call(t *testing.T) {
}

func TestClient_Authorize_integration(t *testing.T) {
	if testing.Short() {
		t.Skip("using -short")
	}

	key := os.Getenv("CLIENTKEY")
	secret := os.Getenv("CLIENTSECRET")
	apiKey := os.Getenv("CLIENTAPIKEY")
	organizationID, _ := strconv.ParseInt(os.Getenv("CLIENTORGANIZATIONID"), 10, 64)

	c := NewClient(key, secret)

	err := c.Authorize(apiKey, int(organizationID))
	if err != nil {
		t.Errorf("Error occured when trying to talk to dinero auth api: %v", err)
	}

	var defaultString string
	if c.token == defaultString {
		t.Errorf("The client did not set the token correctly")
	}

	var defaultInt int
	if c.organizationID == defaultInt {
		t.Errorf("The client did not set the organizationID correctly")
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		clientKey    string
		clientSecret string
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			"test",
			args{"testKey", "testSecret"},
			Client{
				clientKey:    "testKey",
				clientSecret: "testSecret",
				userAgent:    "dinero-go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.args.clientKey, tt.args.clientSecret)
			if c.clientKey != tt.want.clientKey &&
				c.clientSecret != tt.want.clientSecret &&
				c.userAgent != tt.want.userAgent {
				t.Errorf("NewClient Returns incorrect keys, error = %v, wanted = %v", c, tt.want)
			}
		})
	}
}
