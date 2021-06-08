//  Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package main

import (
	"fmt"
	"io/ioutil"
	"path"

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

	uuid := "pqnzigq351j2fqgn" // This is a known ID
	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		halt(4, "Error logging into API: %v", err)
	}

	// ---
	// GetPhoto()
	//
	photo, err := client.V1().GetPhoto(uuid)
	if err != nil {
		halt(3, "Error fetching photo: %v", err)
	}

	// ---
	// UpdatePhoto()
	photo.PhotoTitle = "A really great photo!"
	photo, err = client.V1().UpdatePhoto(photo)
	if err != nil {
		halt(2, "Error updating photo: %v", err)
	}

	// ---
	// GetPhotoDownload()
	file, err := client.V1().GetPhotoDownload(photo.UUID)
	if err != nil {
		halt(2, "Error getting photo download: %v", err)
	}

	for _, f := range photo.Files {
		fileName := fmt.Sprintf("ignore_%s", path.Base(f.FileName))
		logger.Always(fileName)
		ioutil.WriteFile(fileName, file, 0666)
	}

}
