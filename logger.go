package main

import "go.uber.org/zap"

func initLogger(c *Config) *zap.Logger {
	if c.Debug {
		return zap.Must(zap.NewDevelopment())
	}

	return zap.Must(zap.NewProduction())
}
