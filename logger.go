package main

import (
	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"go.uber.org/zap"
)

func initLogger(c *config.Config) *zap.Logger {
	if c.Debug {
		return zap.Must(zap.NewDevelopment())
	}

	return zap.Must(zap.NewProduction())
}
