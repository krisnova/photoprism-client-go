package sampleapp

import (
	"path/filepath"
	"runtime"

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
	CreateCommand  = `pcreate`
	DestroyCommand = `pdestroy`
	LogsCommand    = `plogs`
	StartCommand   = `cd ../sample-app && bash -c "./pstart"`
	StopCommand    = `pstop`
)

func (a *SampleApplication) Start() error {
	logger.Info("Starting Application...")
	script, err := NewScriptFromPath(filepath.Join(PrintWorkingDirectory(), StartCommand))
	if err != nil {
		return err
	}
	return script.Interpret()
}

func (a *SampleApplication) Stop() error {
	logger.Info("Stopping Application...")
	script, err := NewScriptFromPath(filepath.Join(PrintWorkingDirectory(), StopCommand))
	if err != nil {
		return err
	}
	return script.Interpret()
}

func (a *SampleApplication) Create() error {
	logger.Info("Create Application...")
	script, err := NewScriptFromPath(filepath.Join(PrintWorkingDirectory(), CreateCommand))
	if err != nil {
		return err
	}
	return script.Interpret()
}

func (a *SampleApplication) Destroy() error {
	logger.Info("Destroying Application...")
	script, err := NewScriptFromPath(filepath.Join(PrintWorkingDirectory(), DestroyCommand))
	if err != nil {
		return err
	}
	return script.Interpret()
}

func (a *SampleApplication) Logs() error {
	logger.Info("Logging Application...")
	script, err := NewScriptFromPath(filepath.Join(PrintWorkingDirectory(), LogsCommand))
	if err != nil {
		return err
	}
	return script.Interpret()
}

func (a *SampleApplication) GetAuth() photoprism.ClientAuthenticator {
	return nil
}

func PrintWorkingDirectory() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		logger.Info("Unable to PWD")
		return ""
	}
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		logger.Info("Unable to PWD: %v", err)
		return ""
	}
	return dir
}
