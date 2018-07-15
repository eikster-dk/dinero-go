package dinero

import (
	"testing"
)

func TestClient_Call(t *testing.T) {
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
