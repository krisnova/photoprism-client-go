package test

import (
	"testing"

	"github.com/kris-nova/photoprism-client-go"
)

// Trailing slash issue
// https://github.com/kris-nova/photoprism-client-go/issues/2
func TestRegressionIssue2(t *testing.T) {
	testStrings := []string{"localhost/", "localhost///////", "localhost//"}
	goal := "localhost"
	for _, str := range testStrings {
		client := photoprism.New(str)
		if client.ConnectionString() != goal {
			t.Error("Failed to trim suffix / in client connection string")
			t.FailNow()
		}
	}

}
