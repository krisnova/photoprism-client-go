package test

import "testing"

func TestHappyIndex(t *testing.T) {
	err := Client.V1().Index()
	if err != nil {
		t.Errorf("failed indexing: %v", err)
	}
}
