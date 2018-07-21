package dinero

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultAuthEndpoint = "https://authz.dinero.dk/dineroapi/oauth/token"
	defaultAPIEndpoint  = "https://api.dinero.dk"
)

// API is an interface
// that wraps the needed methods for communicating with dinero's api
type API interface {
	Authorize(apiKey string, organizationID int) error
	Call(method, path string, b interface{}, o interface{}) error
}

// Client is a wrapper of the httpCLient with all the needed goodies
// that dinero needs
type Client struct {
	clientKey      string
	clientSecret   string
	userAgent      string
	baseURL        string
	token          string
	organizationID int
	httpClient     *http.Client
}

type authorizedResp struct {
	Token        string `json:"access_token"`
	Type         string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// PaginationResult is general information about a given list and how it's paginated
type PaginationResult struct {
	MaxPageSizeAllowed  int
	PageSize            int
	Result              int
	ResultWithoutFilter int
	Page                int
}

// TimestampResponse is a generic response that is used to return the ID and timestamp from dinero
// whenever a succesful action has occured
type TimestampResponse struct {
	ID        string `json:"guid,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// Authorize will authorize the client
// by preparing the clientKey:clientSecret in base64
// send it off to dinero's auth endpoint to receive a
// token that can be used to interact with the api
func (c *Client) Authorize(apiKey string, organizationID int) error {
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

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var authResp authorizedResp
	if err = json.Unmarshal(bytes, &authResp); err != nil {
		return err
	}

	c.token = authResp.Token
	c.organizationID = organizationID

	return nil
}

// Call is a raw method to interact with the dinero API
// it alters the path to adjust the {organizationID} with the correct organizationID
// it combines your path with the base path and adds the authorization header
// ships it off and unmarshals the request into the o param
// todo: make better error handling
func (c *Client) Call(method, path string, body interface{}, o interface{}) error {
	path = strings.Replace(path, "{organizationID}", strconv.Itoa(c.organizationID), 1)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = c.baseURL + path

	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}

		reader = bytes.NewReader(b)
	}

	req, _ := http.NewRequest(method, path, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {
		sBytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(sBytes))
	}

	if o != nil {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return json.Unmarshal(bytes, o)
	}

	return nil
}

// BuildFieldsQuery returns the field query part of the url
func BuildFieldsQuery(fieldQuery ...string) string {
	var query string
	for i, field := range fieldQuery {
		if i == 0 {
			query = field
		} else {
			query = fmt.Sprintf("%v,%v", query, field)
		}
	}

	return query
}

// NewClient prepares a struct that will be used to communicate with
// Dinero's api
func NewClient(clientKey string, clientSecret string) *Client {
	c := Client{
		clientKey:    clientKey,
		clientSecret: clientSecret,
		userAgent:    "dinero-go",
		baseURL:      defaultAPIEndpoint,
		httpClient:   &http.Client{},
	}

	return &c
}
