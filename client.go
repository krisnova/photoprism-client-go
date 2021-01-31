package photoprism

import "github.com/kris-nova/client-go/api/v1"

type Client struct {
	v1client *api.V1Client
}

type ClientAuthenticator interface {
	getKey() string
	getSecret() string
}

// -- [ ClientAuthLogin ] --

// TODO We probably want to base64 encode this
type ClientAuthLogin struct {
	user string
	pass string
}

func NewClientAuthLogin(user, pass string) ClientAuthenticator {
	return &ClientAuthLogin{
		user: user,
		pass: pass,
	}
}

func (c *ClientAuthLogin) getKey() string {
	return c.user
}
func (c *ClientAuthLogin) getSecret() string {
	return c.pass
}

// -- [ ClientAuthToken ] --

// TODO We probably want to base64 encode this
type ClientAuthToken struct {
	key    string
	secret string
}

func NewClientAuthToken(key, secret string) ClientAuthenticator {
	return &ClientAuthToken{
		key:    key,
		secret: secret,
	}
}

func (c *ClientAuthToken) getKey() string {
	return c.key
}
func (c *ClientAuthToken) getSecret() string {
	return c.secret
}

func New(auth ClientAuthenticator) *Client {
	p := &Client{}
	return p
}

func (c *Client) V1() *api.V1Client {
	return c.v1client
}
