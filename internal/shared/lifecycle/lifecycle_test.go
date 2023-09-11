package lifecycle

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GivenTerminationSignal_WhenListenForShutdown_ThenCallsTheFunctionPassed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)

	shutdownFuncCalled := false
	shutdownFunc := func() {
		shutdownFuncCalled = true
	}

	go ListenForApplicationShutDown(shutdownFunc, shutdownSignal)

	shutdownSignal <- syscall.SIGTERM

	time.Sleep(time.Second)
	assert.True(t, shutdownFuncCalled)
}
