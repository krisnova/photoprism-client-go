package api

import (
	"bytes"
	"encoding/json"
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

// V1Client is used to access the V1 Photoprism API.
type V1Client struct {
	downloadToken string
	token         string
	apihost       *url.URL
	client        http.Client
}

// New will only accept a url.URL so that we know
// all errors have been handled up until this point.
func New(connURL *url.URL, token, downloadToken string) *V1Client {
	return &V1Client{
		client:        http.Client{},
		apihost:       connURL,
		token:         token,
		downloadToken: downloadToken,
	}
}

// V1Response is the master HTTP Response object
// for all transactions with the Photoprism API.
type V1Response struct {
	HTTPResponse *http.Response
	StatusCode   int
	Error        error
	Body         []byte
}

// String is used to represent the body of the response as a string.
func (r *V1Response) String() string {
	if r.Error != nil {
		// Handle errors from the HTTP request first
		return fmt.Sprintf(`{\n"StatusCode":%d,\n"Body":"%s"}`, r.StatusCode, string(r.Body))
	}
	return string(r.Body)
}

// JSON will unmarshal onto whatever interface is passed in.
func (r *V1Response) JSON(i interface{}) error {
	if r.Error != nil {
		// Handle errors from the HTTP request first
		return fmt.Errorf("during HTTP request: %v", r.Error)
	}
	err := json.Unmarshal(r.Body, &i)
	if err != nil {
		return fmt.Errorf("during JSON unmarshal: %v", err)
	}
	return nil
}

// GET is the V1 GET function. By design it will check globally for all non 200
// responses and return an error if a non 200 is encountered.
//
// Error Codes:
//    -1   Unable to create request
func (v1 *V1Client) GET(endpointFormat string, a ...interface{}) *V1Response {
	url := v1.Endpoint(fmt.Sprintf(endpointFormat, a...))
	//logger.Debug("GET [%s]", url)
	response := &V1Response{}
	buffer := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, buffer)
	if err != nil {
		response.StatusCode = -1
		response.Error = fmt.Errorf("unable to create new request: %v", err)
		return response
	}
	req.Header.Set("Content-Type", DefaultContentType)
	req.Header.Set("X-Session-Id", v1.token)
	resp, err := v1.client.Do(req)
	if err != nil {
		response.Error = fmt.Errorf("error while executing request: %v", err)
		return response
	}
	response.StatusCode = resp.StatusCode
	response.HTTPResponse = resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error = fmt.Errorf("unable to read body: %v", err)
		return response
	}
	response.Body = body
	if resp.StatusCode != 200 {
		response.Error = fmt.Errorf("[%d]: %s", resp.StatusCode, body)
		return response
	}
	return response
}

// POST is the V1 POST function. By design it will check globally for all non 200
// responses and return an error if a non 200 is encountered.
// POST will accept a payload.
//
// Error Codes:
//    -1   Unable to create request
//    -2   Unable to write payload
//    -3   Unable to JSON Marshal
func (v1 *V1Client) POST(payload interface{}, endpointFormat string, a ...interface{}) *V1Response {
	url := v1.Endpoint(fmt.Sprintf(endpointFormat, a...))
	//logger.Debug("POST [%s]", url)
	response := &V1Response{}
	jBytes, err := json.Marshal(&payload)
	if err != nil {
		response.StatusCode = -3
		response.Error = fmt.Errorf("unable to marshal JSON: %v", err)
		return response
	}
	buffer := &bytes.Buffer{}
	_, err = buffer.Write(jBytes)
	if err != nil {
		response.StatusCode = -2
		response.Error = fmt.Errorf("unable to write payload: %v", err)
		return response
	}

	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		response.StatusCode = -1
		response.Error = fmt.Errorf("unable to create new request: %v", err)
		return response
	}
	req.Header.Set("Content-Type", DefaultContentType)
	req.Header.Set("X-Session-Id", v1.token)
	resp, err := v1.client.Do(req)
	if err != nil {
		response.Error = fmt.Errorf("error while executing request: %v", err)
		return response
	}
	response.StatusCode = resp.StatusCode
	response.HTTPResponse = resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error = fmt.Errorf("unable to read body: %v", err)
		return response
	}
	response.Body = body
	if resp.StatusCode != 200 {
		response.Error = fmt.Errorf("[%d]: %s", resp.StatusCode, body)
		return response
	}
	return response
}

// PUT is the V1 POST function. By design it will check globally for all non 200
// responses and return an error if a non 200 is encountered.
// PUT will accept a payload.
//
// Error Codes:
//    -1   Unable to create request
//    -2   Unable to write payload
//    -3   Unable to JSON Marshal
func (v1 *V1Client) PUT(payload interface{}, endpointFormat string, a ...interface{}) *V1Response {
	url := v1.Endpoint(fmt.Sprintf(endpointFormat, a...))
	//logger.Debug("POST [%s]", url)
	response := &V1Response{}
	jBytes, err := json.Marshal(&payload)
	if err != nil {
		response.StatusCode = -3
		response.Error = fmt.Errorf("unable to marshal JSON: %v", err)
		return response
	}
	buffer := &bytes.Buffer{}
	_, err = buffer.Write(jBytes)
	if err != nil {
		response.StatusCode = -2
		response.Error = fmt.Errorf("unable to write payload: %v", err)
		return response
	}

	req, err := http.NewRequest("PUT", url, buffer)
	if err != nil {
		response.StatusCode = -1
		response.Error = fmt.Errorf("unable to create new request: %v", err)
		return response
	}
	req.Header.Set("Content-Type", DefaultContentType)
	req.Header.Set("X-Session-Id", v1.token)
	resp, err := v1.client.Do(req)
	if err != nil {
		response.Error = fmt.Errorf("error while executing request: %v", err)
		return response
	}
	response.StatusCode = resp.StatusCode
	response.HTTPResponse = resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error = fmt.Errorf("unable to read body: %v", err)
		return response
	}
	response.Body = body
	if resp.StatusCode != 200 {
		response.Error = fmt.Errorf("[%d]: %s", resp.StatusCode, body)
		return response
	}
	return response
}

// DELETE is the V1 POST function. By design it will check globally for all non 200
// responses and return an error if a non 200 is encountered.
// DELETE will accept a payload.
//
// Error Codes:
//    -1   Unable to create request
//    -2   Unable to write payload
//    -3   Unable to JSON Marshal
func (v1 *V1Client) DELETE(payload interface{}, endpointFormat string, a ...interface{}) *V1Response {
	url := v1.Endpoint(fmt.Sprintf(endpointFormat, a...))
	//logger.Debug("POST [%s]", url)
	response := &V1Response{}
	jBytes, err := json.Marshal(&payload)
	if err != nil {
		response.StatusCode = -3
		response.Error = fmt.Errorf("unable to marshal JSON: %v", err)
		return response
	}
	buffer := &bytes.Buffer{}
	_, err = buffer.Write(jBytes)
	if err != nil {
		response.StatusCode = -2
		response.Error = fmt.Errorf("unable to write payload: %v", err)
		return response
	}

	req, err := http.NewRequest("DELETE", url, buffer)
	if err != nil {
		response.StatusCode = -1
		response.Error = fmt.Errorf("unable to create new request: %v", err)
		return response
	}
	req.Header.Set("Content-Type", DefaultContentType)
	req.Header.Set("X-Session-Id", v1.token)
	resp, err := v1.client.Do(req)
	if err != nil {
		response.Error = fmt.Errorf("error while executing request: %v", err)
		return response
	}
	response.StatusCode = resp.StatusCode
	response.HTTPResponse = resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error = fmt.Errorf("unable to read body: %v", err)
		return response
	}
	response.Body = body
	if resp.StatusCode != 200 {
		response.Error = fmt.Errorf("[%d]: %s", resp.StatusCode, body)
		return response
	}
	return response
}

// Endpoint supports "/api/v1" and "api/v1" like strings
// to generate the string type of a given endpoint based on
// a client
//
// v1client := New("http://localhost:8080", "secret-token")
// v1client.EndpointStr("/api/v1/photos") http://localhost:8080/api/v1/photos/
// v1client.EndpointStr("api/v1/photos") http://localhost:8080/api/v1/photos/
func (v1 *V1Client) Endpoint(str string) string {
	var joined string
	if strings.HasPrefix(str, "/") {
		joined = fmt.Sprintf("%s%s", v1.apihost.String(), str)
	} else {
		joined = fmt.Sprintf("%s/%s", v1.apihost.String(), str)
	}
	return joined
}

// SetToken can be used to set an auth token to use as the X-Session-Id
// for this client
func (v1 *V1Client) SetToken(token string) {
	v1.token = token
}
