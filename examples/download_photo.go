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

func main() {
	user := os.Getenv("PHOTOPRISM_USER")
	if user == "" {
		halt(1, "Missing PHOTOPRISM_USER")
	}
	pass := os.Getenv("PHOTOPRISM_PASS")
	if pass == "" {
		halt(2, "Missing PHOTOPRISM_PASS")
	}
	photoclient := photoprism.New()
	photo, err := photoclient.V1().GetPhoto("123")
	if err != nil {
		halt(3, "Error fetching photo: %v", err)
	}
	fmt.Println(*photo)
}
