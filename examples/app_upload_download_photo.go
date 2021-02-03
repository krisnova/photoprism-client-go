package main

import (
	sampleapp "github.com/kris-nova/client-go/sample-app"
	"github.com/kris-nova/logger"
)

func main() {
	logger.Level = 4
	app := sampleapp.New()
	var err error
	err = app.Start()
	if err != nil {
		logger.Critical(err.Error())
	}
	err = app.Stop()
	if err != nil {
		logger.Critical(err.Error())
	}
}
