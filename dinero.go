package dinero

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultAuthEndpoint = "https://authz.dinero.dk/dineroapi/oauth/token"
	defaultAPIEndpoint  = "https://api.dinero.dk"
)

// API is an interface
// that wraps the needed methods for communicating with dinero's api
type API interface {
	Authorize(apiKey string) error
	Call(method, path string, v interface{}) error
}

// Client is a wrapper of the httpCLient with all the needed goodies
// that dinero needs
type Client struct {
	clientKey    string
	clientSecret string
	userAgent    string
	baseURL      *url.URL
	token        string
	httpClient   *http.Client
}

type authorizedResp struct {
	Token        string `json:"access_token"`
	Type         string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// Authorize will authorize the client
// by preparing the clientKey:clientSecret in base64
// send it off to dinero's auth endpoint to receive a
// token that can be used to interact with the api
func (c *Client) Authorize(apiKey string) error {
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.clientKey, c.clientSecret)))
	data := url.Values{}
	data.Add("grant_type", "password")
	data.Add("scope", "read write")
	data.Add("username", apiKey)
	data.Add("password", apiKey)

	req, err := http.NewRequest(http.MethodPost, defaultAuthEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprint("Basic ", encoded))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)

	var authResp authorizedResp
	if err = decoder.Decode(&authResp); err != nil {
		return err
	}

	c.token = authResp.Token

	return nil
}

// Call is..
func (c *Client) Call(method, path string, v interface{}) error {
	return nil
}

// NewClient prepares a struct that will be used to communicate with
// Dinero's api
func NewClient(clientKey string, clientSecret string) *Client {
	c := Client{
		clientKey:    clientKey,
		clientSecret: clientSecret,
		userAgent:    "dinero-go",
		baseURL:      &url.URL{Path: defaultAPIEndpoint},
		httpClient:   &http.Client{},
	}

	return &c
}
