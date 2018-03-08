package runner

import (
	"os"
	"syscall"
)

// Signals defines the list of signals that will cause the server to shutdown
var Signals = []os.Signal{
	os.Interrupt,
	os.Kill,
	syscall.SIGTERM,
}
