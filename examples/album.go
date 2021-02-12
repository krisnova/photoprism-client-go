package main

import (
	"github.com/kris-nova/logger"
	photoprism "github.com/kris-nova/photoprism-client-go"
	"github.com/kris-nova/photoprism-client-go/api/v1"
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

	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}
	logger.Always("Logged in...")

	album := api.Album{
		AlbumTitle: "NovaAlbum",
	}

	newAlbum, err := client.V1().CreateAlbum(album)
	if err != nil {
		halt(2, "Error creating album: %v", err)
	}

	err = client.V1().DeleteAlbums([]string{newAlbum.AlbumUID})
	if err != nil {
		halt(1, "Error deleting album: %v", err)
	}
}
