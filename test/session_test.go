package test

import (
	"testing"

	photoprism "github.com/kris-nova/client-go"
)

// TestHappyLogin should succeed with the good password "missy"
func TestHappyLogin(t *testing.T) {
	client := photoprism.New(WellKnownSampleAppConnectionString)
	err := client.Auth(photoprism.NewClientAuthLogin(WellKnownUser, WellKnownPass))
	if err != nil {
		t.Errorf("expected login: %v", err)
		t.FailNow()
	}
}

// TestSadLogin should fail with the bad password "charlie"
func TestSadLogin(t *testing.T) {
	client := photoprism.New(WellKnownSampleAppConnectionString)
	err := client.Auth(photoprism.NewClientAuthLogin(WellKnownUser, BadPassword))
	if err == nil {
		t.Errorf("expecting error for known bad password")
		t.FailNow()
	}
}
