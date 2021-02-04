package main

import (
	"fmt"
	"os"

	photoprism "github.com/kris-nova/client-go"
)

func main() {
	uuid := os.Getenv("PHOTOPRISM_UUID")
	if uuid == "" {
		halt(2, "Missing PHOTOPRISM_UUID")
	}
	client := photoprism.New(auth())
	photo, err := client.V1().GetPhoto(uuid)
	if err != nil {
		halt(3, "Error fetching photo: %v", err)
	}

	fmt.Println(*photo)
}
