package test

import (
	"os"
	"strings"
	"testing"

	photoprism "github.com/kris-nova/client-go"

	"github.com/kris-nova/logger"

	sampleapp "github.com/kris-nova/client-go/sample-app"
)

const (
	WellKnownUser                      = "admin"
	WellKnownPass                      = "missy"
	BadPassword                        = "charlie"
	WellKnownPhotoID                   = "pqnzigq351j2fqgn" // This is a photo in the persistent sample app
	UnknownPhotoID                     = "1234567890"
	WellKnownSampleAppConnectionString = "http://localhost:8080"
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
	defer func() {
		//err := app.Stop()
		//if err != nil {
		//	logger.Critical("Failure stopping application: %v", err)
		//	os.Exit(100)
		//}
		logger.Always("Success!")
		os.Exit(0)
	}()
	if err != nil {
		logger.Critical("Unable to start app: %v", err)
		os.Exit(2)
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
