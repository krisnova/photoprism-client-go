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
package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"

	photoprism "github.com/drummonds/photoprism-client-go"

	"github.com/kris-nova/logger"

	"github.com/drummonds/photoprism-client-go/api/v1"
	sampleapp "github.com/drummonds/photoprism-client-go/sample-app"
)

const (
	WellKnownUser                      = "admin"
	WellKnownPass                      = "missy"
	BadPassword                        = "charlie"
	UnknownPhotoID                     = "1234567890"
	UnknownAlbumID                     = "1234567890"
	WellKnownSampleAppConnectionString = "http://localhost:8080"
	UnknownCategory                    = "Furries"
	WellKnownAlbumTitle                = "TestAlbum"
)

// Dynamically created test variables for duration of test
var (
	WellKnownPhotoID string
	WellKnownAlbumID string
)

// Client is a pre-authenticated client that can be used
// internally to access the SDK
var Client *photoprism.Client

func cleanStorageDirectory(app *sampleapp.SampleApplication) {
	// Close down any old env (to make sure clean start)
	_ = app.Stop()
	// ignore error as may or may not be running
	_ = app.Destroy()
	// ignore error as may or may have been created but it isn't any longer
	//This is going to delete the storage directory and then build ready for a
	// new create of photoprism copying in some template files
	// the aim is to have a set of minimal template files needed to build an environment
	_, filename, _, _ := runtime.Caller(1)
	storagePath := path.Join(path.Dir(filename), "../sample-app/photoprism/storage")
	cmdStr := fmt.Sprintf("cd %s; cd ../sample-app/photoprism; sudo rm -rf storage",
		path.Dir(filename))
	// Need to use sudo to remove directories as using sudo to install
	_, err := exec.Command("bash", "-c", cmdStr).Output()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(storagePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateWellKnownAlbum() {
	// Find a photo
	result, err := Client.V1().GetPhotos(&api.PhotoOptions{
		Count: 2,
		// AlbumUID: WellKnownAlbumID,
	})
	if err != nil {
		logger.Critical(fmt.Sprintf("expected success getting well known photo: %v, %v", err, result))
	}

	// need well known album
	album := api.Album{
		AlbumTitle: WellKnownAlbumTitle,
	}

	wellKnownAlbum, err := Client.V1().CreateAlbum(album)
	if err != nil {
		logger.Critical(fmt.Sprintf("expected to create album: %v", err))
	}
	WellKnownAlbumID = wellKnownAlbum.AlbumUID
	// Add Photos to album
	photos := []string{
		result[0].PhotoUID,
		result[1].PhotoUID,
	}
	err = Client.V1().AddPhotosToAlbum(wellKnownAlbum.AlbumUID, photos)
	if err != nil {
		logger.Critical(fmt.Sprintf("expected to add photos to album: %v", err))
	}
}

func TestMain(m *testing.M) {
	logger.Level = 4
	app := sampleapp.New()
	cleanStorageDirectory(app)
	err := app.Create()
	// This System will create a new cluster
	// if needed. Otherwise it will log the
	// error at the INFO level.
	if err != nil {
		if !strings.Contains(err.Error(), "The container name \"/photoprism\" is already in use") {
			logger.Critical("Unable to create app: %v", err)
			os.Exit(1)
		}
		///logger.Debug(err.Error())
	}
	err = app.Start()
	if false { // Turn off if you want to see what state the sample app was in
		logger.Always("Clean up")
		cleanStorageDirectory(app)
	}
	if err == nil {
		logger.Always("Stopping Photoprism Sample App for Unit Tests (deferred)")
		defer func() {
			err := app.Stop()
			if err != nil {
				logger.Critical("Failure stopping application: %v", err)
				os.Exit(100)
			}
			logger.Always("Success!")
			os.Exit(0)
		}()
	} else {
		logger.Always("Photoprism already running...")
	}

	// --- [ Client ] ---
	client := photoprism.New(WellKnownSampleAppConnectionString)
	err = client.Auth(photoprism.NewClientAuthLogin(WellKnownUser, WellKnownPass))
	if err != nil {
		logger.Critical("Error during testing auth: %v", err)
		os.Exit(3)
	}
	Client = client

	// --- [ Prep photoprism ]
	// As starting un-indexed need to index and create well known albums
	err = Client.V1().Index()
	if err != nil {
		logger.Critical(fmt.Sprintf("expected success index original photos: %v", err))
	}
	// Find a photo
	result, err := Client.V1().GetPhotos(&api.PhotoOptions{
		Count: 2,
		// AlbumUID: WellKnownAlbumID,
	})
	if err != nil {
		logger.Critical(fmt.Sprintf("expected success getting well known photo: %v, %v", err, result))
	}
	WellKnownPhotoID = result[0].PhotoUID
	CreateWellKnownAlbum()

	// --- [ Tests ] ----
	exitCode := m.Run()
	if exitCode != 0 {
		logger.Critical("Failure!")
		os.Exit(100)
	}
	// --- [ Tests ] ---

}
