package main

import (
	photoprism "github.com/kris-nova/client-go"
	"github.com/kris-nova/logger"
)

func main() {
	logger.Level = 4
	creds := photoprism.NewClientAuthLogin("admin", "missy")
	client := photoprism.New("localhost:8080")
	err := client.Auth(creds)
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Login Success!")
}
