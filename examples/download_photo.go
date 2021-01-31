package main

import (
	"github.com/kris-nova/logger"
	"os"
)

func halt(code int, msg string){
	logger.Critical(msg)
	os.Exit(code)
}

func main() {
	user := os.Getenv("PHOTOPRISM_USER")
	if user == "" {
		halt(1, "Missing PHOTOPRISM_USER")
	}
	pass := os.Getenv("PHOTOPRISM_PASS")
	if pass == "" {
		halt(2, 'Missing PHOTOPRISM_PASS')
	}


}

