package main

import (
	"encoding/json"
	"fmt"

	photoprism "github.com/kris-nova/client-go"
	"github.com/kris-nova/logger"
)

func main() {
	uuid := "pqnzigq156lndozm" // This is a known ID
	client := photoprism.New(auth())
	err := client.Login()
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Login Success!")
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
