// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"io/ioutil"
	"path"

	photoprism "github.com/drummonds/photoprism-client-go"
	"github.com/drummonds/photoprism-client-go/api/v1"
	common "github.com/drummonds/photoprism-client-go/examples/common"
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

	client := photoprism.New("http://localhost:8080")
	err := client.Auth(photoprism.NewClientAuthLogin("admin", "missy"))
	if err != nil {
		common.Halt(4, "Error logging into API: %v", err)
	}

	result, err := client.V1().GetPhotos(&api.PhotoOptions{
		Count: 2,
		// AlbumUID: WellKnownAlbumID,
	})
	if err != nil {
		common.Halt(3, "expected success getting well known photo: %v, %v", err, result)
	}
	uuid := result[0].PhotoUID

	// ---
	// GetPhoto()
	//
	photo, err := client.V1().GetPhoto(uuid)
	if err != nil {
		common.Halt(3, "Error fetching photo: %v", err)
	}

	// ---
	// UpdatePhoto()
	photo.PhotoTitle = "A really great photo!"
	photo, err = client.V1().UpdatePhoto(photo)
	if err != nil {
		common.Halt(2, "Error updating photo: %v", err)
	}

	// ---
	// GetPhotoDownload()
	file, err := client.V1().GetPhotoDownload(photo.UUID)
	if err != nil {
		common.Halt(2, "Error getting photo download: %v", err)
	}

	for _, f := range photo.Files {
		fileName := fmt.Sprintf("ignore_%s", path.Base(f.FileName))
		logger.Always(fileName)
		ioutil.WriteFile(fileName, file, 0666)
	}

}
