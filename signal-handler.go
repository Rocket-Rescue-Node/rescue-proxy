package main

import (
	"os"
	"os/signal"
	"syscall"
)

func handleSignals(signals ...os.Signal) chan bool {

	exit := make(chan bool)
	c := make(chan os.Signal, 1)

	// Always wait for SIGTERM at a minimum
	signal.Notify(c, syscall.SIGTERM)

	if len(signals) != 0 {
		for _, s := range signals {
			if s == syscall.SIGTERM {
				continue
			}
			signal.Notify(c, s)
		}
	}

	go func() {
		// Block until signal is received
		<-c

		// Allow subsequent signals to quickly shut down by removing the trap
		signal.Reset()

		// Close the channels
		close(c)
		close(exit)
	}()

	return exit
}
