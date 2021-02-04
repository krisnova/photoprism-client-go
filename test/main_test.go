package test

import (
	"os"
	"strings"
	"testing"

	"github.com/kris-nova/logger"

	sampleapp "github.com/kris-nova/client-go/sample-app"
)

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

	// --- [ Tests ] ----
	exitCode := m.Run()
	if exitCode != 0 {
		logger.Critical("Failure!")
		os.Exit(100)
	}
	// --- [ Tests ] ---

}

func TestHappyAPI(t *testing.T) {
	// Code to validate the API
	if true {
		t.Logf("Success!\n")
	} else {
		t.Errorf("Failure!\n")
	}
}

func TestSadAPI(t *testing.T) {
	// Code to validate the API
	if !false {
		t.Logf("Success!\n")
	} else {
		t.Errorf("Failure!\n")
	}
}
