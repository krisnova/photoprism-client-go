package sampleapp

import (
	"github.com/kris-nova/logger"

	photoprism "github.com/kris-nova/client-go"
)

type SampleApplication struct {
}

func New() *SampleApplication {
	app := &SampleApplication{}
	return app
}

// These are the bash scripts that can be used
// to start/stop the Photoprism test application
var (
	CreateCommand  = `./pcreate`
	DestroyCommand = `./pdestroy`
	LogsCommand    = `./plogs`
	StartCommand   = `./pstart`
	StopCommand    = `./pstop`
)

func (a *SampleApplication) Start() error {
	logger.Info("Starting Application...")
	script := NewScript(StartCommand)
	return script.Interpret()
}

func (a *SampleApplication) Stop() error {
	logger.Info("Stopping Application...")
	script := NewScript(StopCommand)
	return script.Interpret()
}

func (a *SampleApplication) GetAuth() photoprism.ClientAuthenticator {
	return nil
}
