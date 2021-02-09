package main

import (
	"encoding/json"
	"fmt"

	photoprism "github.com/kris-nova/client-go"
	"github.com/kris-nova/logger"
)

func main() {
	logger.Level = 4
	uuid := "pqnzigq351j2fqgn" // This is a known ID
	creds := photoprism.NewClientAuthLogin("admin", "missy")
	client := photoprism.New("localhost:8080")
	err := client.Auth(creds)
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	//logger.Always("Login Success!")
	photo, err := client.V1().GetPhoto(uuid)
	if err != nil {
		halt(3, "Error fetching photo: %v", err)
	}
	bytes, err := json.Marshal(photo)
	if err != nil {
		halt(5, "Error: %v", err)
	}
	fmt.Println(string(bytes))
}
