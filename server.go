package runner

import (
	"os"
	"os/signal"
)

// Server defines a runable interface. The Run() method should not block (any
// background tasks should be run in goroutines)
type Server interface {
	Run() error
	Stop() error
}

// RunServer starts up a Server, then blocks until a signal is received to shut
// down at which time the Server's Stop() method is called
func RunServer(server Server) error {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, Signals...)

	err := server.Run()
	if err != nil {
		return err
	}
	<-stopChan // wait for signals
	return server.Stop()
}
