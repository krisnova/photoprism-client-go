package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// DefaultContentType is the content type header the API expects
	DefaultContentType string = "application/json; charset=utf-8"
)

type V1Client struct {
	token            string
	connectionString string
	client           http.Client
}

func New(connection, token string) *V1Client {
	c := http.Client{}
	return &V1Client{
		client:           c,
		connectionString: connection,
		token:            token,
	}
}

func (v1 *V1Client) SetConnectionString(connection string) {
	v1.connectionString = connection
}

func (v1 *V1Client) SetToken(token string) {
	v1.token = token
}

// GET is the V1 GET function. By design it will check globally for all non 200
// responses and return an error if a non 200 is encountered.
func (v1 *V1Client) GET(format string, a ...interface{}) (*http.Response, error) {
	str := fmt.Sprintf(format, a...)
	//logger.Debug("GET [%s]", str)
	url := v1.EndpointStr(str)
	buffer := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, buffer)
	if err != nil {
		return nil, fmt.Errorf("unable to generate new request: %v", err)
	}
	req.Header.Set("Content-Type", DefaultContentType)
	req.Header.Set("X-Session-Id", v1.token)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, fmt.Errorf("[%d]: unable to read body: %v", err)
		}
		return resp, fmt.Errorf("[%d]: %s", resp.StatusCode, body)
	}
	return resp, nil
}

func (v1 *V1Client) EndpointStr(str string) string {
	if strings.HasPrefix("/", str) {
		str = fmt.Sprintf("%s%s", v1.connectionString, str)
	} else {
		str = fmt.Sprintf("%s/%s", v1.connectionString, str)
	}
	return str
}

func (v1 *V1Client) EndpointURL(str string) (*url.URL, error) {
	if strings.HasPrefix("/", str) {
		str = fmt.Sprintf("%s%s", v1.connectionString, str)
	} else {
		str = fmt.Sprintf("%s/%s", v1.connectionString, str)
	}
	return url.Parse(str)
}
