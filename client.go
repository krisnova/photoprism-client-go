package photoprism

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	v1 "github.com/kris-nova/photoprism-client-go/api/v1"
)

const (
	// DefaultContentType is the content type header the API expects
	APIContentType string = "application/json; charset=utf-8"

	// Default Host Configuration
	APIAuthHeaderKey string = "X-Session-Id"
)

// New is used to create a new Client to authenticate with
// Photoprism.
func New(connectionString string) *Client {
	for strings.HasSuffix(connectionString, "/") {
		connectionString = connectionString[:len(connectionString)-1]
	}

	c := &Client{
		contentType:      APIContentType,
		connectionString: connectionString,
	}
	return c
}

func (c *Client) ConnectionString() string {
	return c.connectionString
}

// Client represents a client to a Photoprism application
type Client struct {
	v1client         *v1.V1Client
	authenticator    ClientAuthenticator
	contentType      string
	connectionString string
	connectionURL    *url.URL
}

// ClientAuthenticator is used to store the secret
// data for authenticating with the Photoprism API
type ClientAuthenticator interface {
	getKey() string
	getSecret() string
	JSON() ([]byte, error)
}

// ClientAuthLogin holds secret login information
type ClientAuthLogin struct {
	authPayload
}

type authPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClientAuthLogin is used to build a new login struct
func NewClientAuthLogin(user, pass string) ClientAuthenticator {
	return &ClientAuthLogin{
		authPayload: authPayload{
			Username: user,
			Password: pass,
		},
	}
}

// getKey is used internally to get the key with any modifiers
func (c *ClientAuthLogin) getKey() string {
	return c.authPayload.Username
}

// getKey is used internally to get the secret with any modifiers
func (c *ClientAuthLogin) getSecret() string {
	return c.authPayload.Password
}

// JSON is used to marshal the fields to JSON
func (c *ClientAuthLogin) JSON() ([]byte, error) {
	return json.Marshal(c)
}

// V1 is used to access the V1 version of the Photoprism API
func (c *Client) V1() *v1.V1Client {
	return c.v1client
}

// Login is used to attempt to authenticate with the Photoprism API
func (c *Client) Auth(auth ClientAuthenticator) error {
	c.authenticator = auth

	// @kris-nova We are returning V1 by default
	return c.LoginV1()
}

// POST /api/v1/session
//
// Data: {username: "admin", password: "missy"}
func (c *Client) LoginV1() error {
	// Auth wil also validate the connection string
	// We do this here so that New() will never return
	// an error.
	url, err := url.Parse(c.connectionString)
	if err != nil {
		return fmt.Errorf("unable to parse connection string url [%s]: %v", c.connectionString, err)
	}
	c.connectionURL = url

	body, err := c.authenticator.JSON()
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}
	buffer := bytes.NewBuffer(body)
	resp, err := http.Post(c.Endpoint("api/v1/session"), c.contentType, buffer)
	if err != nil {
		return fmt.Errorf("authentication error: %v", err)
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to parse body for [%d] response: %v", resp.StatusCode, err)
		}
		return fmt.Errorf("login error [%d] %s", resp.StatusCode, body)
	}

	// --- JSON Auth Response on to Options ---
	cfg := &Config{
		Config: &Options{},
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to parse auth body: %v", err)
	}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return fmt.Errorf("unable to json unmarshal auth body: %v", err)
	}

	token := resp.Header.Get(APIAuthHeaderKey)
	if token == "" {
		return fmt.Errorf("missing auth token from successful login")
	}
	c.v1client = v1.New(c.connectionURL, token, cfg.Config.DownloadToken)
	return nil
}

// Endpoint is used to calculate a FQN for a given API endpoint
// based on the API version and Host/Port
func (c *Client) Endpoint(str string) string {
	if strings.HasPrefix("/", str) {
		str = fmt.Sprintf("%s%s", c.connectionString, str)
	} else {
		str = fmt.Sprintf("%s/%s", c.connectionString, str)
	}
	//logger.Debug(str)
	return str
}
