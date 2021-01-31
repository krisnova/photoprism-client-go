package api

import "net/http"

type Meta struct {
	requested bool
	response  *http.Response
}

// TODO We need an http.Client and an object interface{}
func (m *Meta) Request() error {
	m.requested = true
	return nil
}

func (m *Meta) Response() *http.Response {
	if !m.requested {
		return nil
	}
	return m.response
}
