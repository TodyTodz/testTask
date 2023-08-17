package app_utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	. "github.com/NGRsoftlab/ngr-logging"
)

func notifyChan(in chan os.Signal) {
	signal.Notify(in, os.Interrupt, syscall.SIGTERM)
	signal.Notify(in, os.Interrupt, syscall.SIGINT)
}

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

func WaitSignalsContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		Logger.Info("Waiting finished")
	}

	return nil
}

