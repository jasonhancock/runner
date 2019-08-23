package runner

import (
	"context"
	"os"
	"os/signal"
)

// Server defines a runable interface. The Run() method should not block (any
// background tasks should be run in goroutines)
type Server interface {
	Run() error
	Stop() error
}

// Context returns a context that will be cancelled when a signal is received to shut down.
func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, Signals...)

	go func() {
		<-stopChan
		cancel()
	}()

	return ctx
}

// RunServer starts up a Server, then blocks until a signal is received to shut
// down at which time the Server's Stop() method is called
func RunServer(server Server) error {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, Signals...)
	if err := server.Run(); err != nil {
		return err
	}
	<-stopChan // wait for signals
	return server.Stop()
}
