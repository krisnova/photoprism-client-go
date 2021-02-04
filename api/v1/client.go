package api

import (
	"fmt"
	"net/http"
	"strings"
)

type V1Client struct {
	token            string
	connectionString string
}

func (v1 *V1Client) SetConnectionString(connection string) {
	v1.connectionString = connection
}

func (v1 *V1Client) SetToken(token string) {
	v1.token = token
}

func (v1 *V1Client) GET(format string, a ...interface{}) (*http.Response, error) {
	url := v1.Endpoint(fmt.Sprintf(format, a))
	return http.Get(url)
}

func (v1 *V1Client) Endpoint(str string) string {
	if strings.HasPrefix("/", str) {
		return fmt.Sprintf("%s%s", v1.connectionString, str)
	}
	return fmt.Sprintf("%s/%s", v1.connectionString, str)
}
