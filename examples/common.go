package main

import (
	"fmt"
	"os"

	photoprism "github.com/kris-nova/client-go"

	"github.com/kris-nova/logger"
)

// format string, a ...interface{}
func halt(code int, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	logger.Critical(str)
	os.Exit(code)
}

func auth() photoprism.ClientAuthenticator {
	user := os.Getenv("PHOTOPRISM_USER")
	if user == "" {
		halt(1, "Missing PHOTOPRISM_USER")
	}
	pass := os.Getenv("PHOTOPRISM_PASS")
	if pass == "" {
		halt(2, "Missing PHOTOPRISM_PASS")
	}
	auth := photoprism.NewClientAuthLogin(user, pass)
	return auth
}
