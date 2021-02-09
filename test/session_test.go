package test

import (
	"testing"

	photoprism "github.com/kris-nova/client-go"
)

const (
	WellKnownUser = "admin"
	WellKnownPass = "missy"
	BadPassword   = "charlie"
)

// TestHappyLogin should succeed with the good password "missy"
func TestHappyLogin(t *testing.T) {
	creds := photoprism.NewClientAuthLogin(WellKnownUser, WellKnownPass)
	client := photoprism.New("localhost:8080")
	err := client.Auth(creds)
	if err != nil {
		t.Errorf("invalid login: %v", err)
	}
}

// TestSadLogin should fail with the bad password "charlie"
func TestSadLogin(t *testing.T) {
	creds := photoprism.NewClientAuthLogin(WellKnownUser, BadPassword)
	client := photoprism.New("localhost:8080")
	err := client.Auth(creds)
	if err == nil {
		t.Errorf("Missing error for known bad password")
		return
	}
	t.Logf("Successful bad password auth attempt: %v", err)
}
