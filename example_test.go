package runner

import "log"

type myServer struct {}

func (s *myServer) Run() error {
    log.Println("starting...")
    return nil
}

func (s *myServer) Stop() error {
    log.Println("stopping...")
    return nil
}

func Example() {
    // create a context that will be cancelled when an os signal tells the process to shut down.
    ctx := Context()

    // Fire up some background job that will shutdown when ctx is cancelled.
    go func() {
        // some background job....
        <-ctx.Done()
    }()

    // start up a server. This call will block.
    RunServer(&myServer{})
}
