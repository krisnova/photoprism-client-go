package main

import (
	"fmt"
	photoprism "github.com/kris-nova/client-go"
	"github.com/kris-nova/logger"
)

func main() {
	// ---
	// Log Level 4 (Most)
	// Log Level 3
	// Log Level 2
	// Log Level 1
	// Log Level 0 (Least)
	//
	logger.Level = 4
	//
	// ---
	uuid := "aqnzih81icziiyae"

	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Logged in...")

	album, err := client.V1().GetAlbum(uuid)
	fmt.Println(album)

}
