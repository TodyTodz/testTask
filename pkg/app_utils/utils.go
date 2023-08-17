//Package for app utils:
// - notify additional channel for kill signal
// - waiting for signals with context and channel
package app_utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	. "github.com/NGRsoftlab/ngr-logging"
)

//Notify additional channel for kill signals
func notifyChan(in chan os.Signal) {
	signal.Notify(in, os.Interrupt, syscall.SIGTERM)
	signal.Notify(in, os.Interrupt, syscall.SIGINT)
}

//Waiting for any of signals
func WaitSignals(ctx context.Context) error {
	sig := make(chan os.Signal, 1)
	notifyChan(sig)

	select {
	case <-ctx.Done():
	case <-sig:
		Logger.Info("Waiting finished")
	}

	return WaitSignalsContext(ctx)
}

//Waiting for contest kill signals
// Here able to be additional wrappers for shutdown the service
func WaitSignalsContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		Logger.Info("Waiting finished")
	}

	return nil
}

