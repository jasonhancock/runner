package runner

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRunner(t *testing.T) {
	var run, stop int
	newServer := func() *ServerMock {
		return &ServerMock{
			RunFunc: func() error {
				run++
				return nil
			},
			StopFunc: func() error {
				stop++
				return nil
			},
		}
	}

	ctx := Context()

	go RunServer(newServer())
	go RunServer(newServer())
	go RunServer(newServer())

	go func() {
		<-ctx.Done()
		stop++
	}()

	time.Sleep(100 * time.Millisecond)

	require.Equal(t, 3, run)

	// we fake the signal by closing the stop channel
	close(stopChan)
	time.Sleep(100 * time.Millisecond)

	require.Equal(t, 4, stop)
}
