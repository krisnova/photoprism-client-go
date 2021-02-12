package main

import (
	"fmt"

	"github.com/kris-nova/logger"
	photoprism "github.com/kris-nova/photoprism-client-go"
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
	if err != nil {
		halt(3, "Error getting album %s", uuid)
	}
	fmt.Println(album)

	albums, err := client.V1().GetAlbums(nil)
	if err != nil {
		halt(2, "Error listing albums: %v", err)
	}
	for _, album := range albums {
		fmt.Println(album)
	}

}
