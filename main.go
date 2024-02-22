package main

import (
	"context"
	"os"

	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"go.uber.org/zap"
)

func loop(signalChan chan bool, errs chan error) error {
	for {
		select {
		case <-signalChan:
			return nil
		case err := <-errs:
			return err
		}
	}
}

func main() {

	// Initialize config
	config := config.InitFlags()

	// Create the service
	service := NewService(config)

	logger := initLogger(config)
	defer func() {
		logger.Info("rescue-proxy shutdown completed, flushing logs")
		_ = logger.Sync()
	}()

	service.Logger = logger

	// Trap signals
	logger.Debug("Trapping SIGTERM and SIGINT")
	signalChan := handleSignals(os.Interrupt)

	errs := service.Run(context.Background())

	if err := loop(signalChan, errs); err != nil {
		logger.Panic("error running service", zap.Error(err))
	}

	// Shut down gracefully
	logger.Info("Received signal, shutting down")
	if err := service.Stop(); err != nil {
		logger.Panic("error stopping service", zap.Error(err))
	}

}
