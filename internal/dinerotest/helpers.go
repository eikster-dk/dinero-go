package dinerotest

import (
	"os"
	"strconv"
)

// Text to show when running go test -short
const IntegrationTestText = "Skipped integration test by using short"

// common way to get environment variables for doing integration tests
func GetClientKeysForIntegrationTesting() (string, string, string, int) {
	key := os.Getenv("CLIENTKEY")
	secret := os.Getenv("CLIENTSECRET")
	apiKey := os.Getenv("CLIENTAPIKEY")
	organizationID, _ := strconv.ParseInt(os.Getenv("CLIENTORGANIZATIONID"), 10, 64)

	return key, secret, apiKey, int(organizationID)
}
