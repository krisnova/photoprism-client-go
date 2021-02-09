package photoprism

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kris-nova/client-go/api/v1"
)

const (
	// DefaultContentType is the content type header the API expects
	DefaultContentType string = "application/json; charset=utf-8"

	// Default Host Configuration
	DefaultHost             string = "localhost"
	DefaultLoopback         string = "127.0.0.1"
	DefaultPort             string = "8080"
	DefaultConnectionString string = "http://localhost:8080"
	DefaultTokenKey         string = "X-Session-Id"
)

func New(connectionString string) *Client {
	c := &Client{
		contentType:      DefaultContentType,
		connectionString: DefaultConnectionString,
	}
	return c
}

// Client represents a client to a Photoprism application
type Client struct {
	v1client         *api.V1Client
	authenticator    ClientAuthenticator
	contentType      string
	connectionString string
}

// ClientAuthenticator is used to store the secret
// data for authenticating with the Photoprism API
//
// TODO @kris-nova obfuscate the data, and make immutable and unexported fields
type ClientAuthenticator interface {
	getKey() string
	getSecret() string
	JSON() ([]byte, error)
}

// -- [ ClientAuthLogin ] --

type ClientAuthLogin struct {
	// Request
	Username string `json:"username"`
	Password string `json:"password"`

	// Response
}

// NewClientAuthLogin is used to build a new login struct
func NewClientAuthLogin(user, pass string) ClientAuthenticator {
	return &ClientAuthLogin{
		Username: user,
		Password: pass,
	}
}

func (c *ClientAuthLogin) getKey() string {
	return c.Username
}
func (c *ClientAuthLogin) getSecret() string {
	return c.Password
}

// JSON is used to marshal the fields to JSON
func (c *ClientAuthLogin) JSON() ([]byte, error) {
	return json.Marshal(c)
}

// V1 is used to access the V1 version of the Photoprism API
func (c *Client) V1() *api.V1Client {
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
	token := resp.Header.Get(DefaultTokenKey)
	if token == "" {
		return fmt.Errorf("missing auth token from successful login")
	}
	c.v1client = api.New(c.connectionString, token)
	return nil
}

// Endpoint is used to calculate a FQN for a given API endpoint
// based on the API version and Host/Port
func (c *Client) Endpoint(str string) string {
	if strings.HasPrefix("/", str) {
		return fmt.Sprintf("%s%s", c.connectionString, str)
	}
	return fmt.Sprintf("%s/%s", c.connectionString, str)
}

// ----------------------------------------------------------------------------------------
// Dump from Chrome/Network

// [REQUEST]
//Request URL: http://localhost:8080/api/v1/session
//Request Method: POST
//Status Code: 200 OK
//Remote Address: 127.0.0.1:8080
//Referrer Policy: strict-origin-when-cross-origin

// [RESPONSE HEADERS]
//Content-Type: application/json; charset=utf-8
//Date: Thu, 04 Feb 2021 03:27:03 GMT
//Transfer-Encoding: chunked
//X-Session-Id: d92837cb1c41e37b9993d25e282efb3b337b6ae609a687d9

// [REQUEST HEADERS]
//Accept: application/json, text/plain, */*
//Accept-Encoding: gzip, deflate, br
//Accept-Language: en-US,en;q=0.9
//Connection: keep-alive
//Content-Length: 39
//Content-Type: application/json;charset=UTF-8
//Host: localhost:8080
//Origin: http://localhost:8080
//Referer: http://localhost:8080/login
//Sec-Fetch-Dest: empty
//Sec-Fetch-Mode: cors
//Sec-Fetch-Site: same-origin
//User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36
//X-Client-Hash: 2607a5a5
//X-Client-Version: 210121-07e559df-Linux-x86_64

// [POST DATA]
//{username: "admin", password: "missy"}
//password: "missy"
//username: "admin"
