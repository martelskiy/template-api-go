package lifecycle

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/martelskiy/template-api-go/internal/shared/logger"
)

var log = logger.Get()

func ListenForApplicationShutDown(shutdownFunc func(), signalChannel chan os.Signal) {
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChannel
	switch sig {
	case os.Interrupt, syscall.SIGTERM:
		log.Info("shutdown signal received")
		shutdownFunc()
	}
}

func StopApplication(message string) {
	log.With(message).Error("stopping application")
	os.Exit(-1)
}
