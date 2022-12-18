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
package sampleapp

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kris-nova/logger"

	"github.com/drummonds/photoprism-client-go"
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
	StartCommand   = `pstart`
	StopCommand    = `pstop`
)

func (a *SampleApplication) Start() error {
	logger.Info("Starting Application...")
	cmdStr := filepath.Join(PrintWorkingDirectory(), StartCommand)
	result, err := Exec(cmdStr)
	if err != nil {
		return err
	}
	logger.Debug(result.Stdout())
	timeout := 15
	logger.Info(fmt.Sprintf("Wait till running ca %v secs", timeout))
	time.Sleep(time.Second * time.Duration(timeout))
	logger.Info("Hope now working, 5 secs was not enough")
	return nil
}

func (a *SampleApplication) Stop() error {
	logger.Info("Stopping Application...")
	cmdStr := filepath.Join(PrintWorkingDirectory(), StopCommand)
	_, err := Exec(cmdStr)
	return err
}

func (a *SampleApplication) Create() error {
	logger.Info("Create Application...")
	cmdStr := filepath.Join(PrintWorkingDirectory(), CreateCommand)
	result, err := Exec(cmdStr)
	logger.Info("From create application:\n%s\n", result.stdout)
	return err
}

func (a *SampleApplication) Destroy() error {
	logger.Info("Destroying Application...")
	cmdStr := filepath.Join(PrintWorkingDirectory(), DestroyCommand)
	_, err := Exec(cmdStr)
	return err
}

func (a *SampleApplication) Logs() error {
	logger.Info("Logging Application...")
	cmdStr := filepath.Join(PrintWorkingDirectory(), LogsCommand)
	_, err := Exec(cmdStr)
	return err
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
