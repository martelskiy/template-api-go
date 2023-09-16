package lifecycle

import (
	"os"
	"sync"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenTerminationSignal_WhenListenForShutdown_ThenCallsTheFunctionPassed(t *testing.T) {
	shutdownSignal := make(chan os.Signal, 1)
	var wg sync.WaitGroup
	shutdownFuncCalled := false
	shutdownFunc := func() {
		defer wg.Done()
		shutdownFuncCalled = true
	}
	wg.Add(1)

	go ListenForApplicationShutDown(shutdownFunc, shutdownSignal)

	shutdownSignal <- syscall.SIGTERM
	wg.Wait()

	assert.True(t, shutdownFuncCalled)
}
