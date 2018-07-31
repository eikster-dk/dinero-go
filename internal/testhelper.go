package internal

import (
	"github.com/eikc/dinero-go"
	"github.com/eikc/dinero-go/internal/dinerotest"
)

// GetClient gets and  authorize the client with environment variables.
func GetClient() *dinero.Client {
	key, secret, apiKey, organizationID := dinerotest.GetClientKeysForIntegrationTesting()

	c := dinero.New(key, secret)
	c.Authorize(apiKey, organizationID)

	return c
}
