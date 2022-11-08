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
	"os"
	"strings"
	"testing"

	photoprism "github.com/drummonds/photoprism-client-go"

	"github.com/kris-nova/logger"

	sampleapp "github.com/drummonds/photoprism-client-go/sample-app"
)

const (
	WellKnownUser                      = "admin"
	WellKnownPass                      = "missy"
	BadPassword                        = "charlie"
	WellKnownPhotoID                   = "pqnzigq351j2fqgn" // This is a photo in the persistent sample app
	UnknownPhotoID                     = "1234567890"
	WellKnownAlbumID                   = "aqnzih81icziiyae"
	UnknownAlbumID                     = "1234567890"
	WellKnownSampleAppConnectionString = "http://localhost:8080"
	UnknownCategory                    = "Furries"
	WellKnownAlbumTitle                = "TestAlbum"
)

// Client is a pre-authenticated client that can be used
// internally to access the SDK
var Client *photoprism.Client

func TestMain(m *testing.M) {
	logger.Level = 4
	app := sampleapp.New()
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
	if err == nil {
		logger.Always("Stopping Photoprism Sample App for Unit Tests")
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

	// --- [ Tests ] ----
	exitCode := m.Run()
	if exitCode != 0 {
		logger.Critical("Failure!")
		os.Exit(100)
	}
	// --- [ Tests ] ---

}
