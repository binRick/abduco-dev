package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx context.Context
)

func init() {
	ctx = cancelContext()
}

func cancelContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-s
		if false {
			fmt.Fprintf(os.Stderr, "\n\ncancelContext> TRIGGER\n\n\n")
		}
		cancel()
	}()
	if false {
		fmt.Fprintf(os.Stderr, "\n\ncancelContext> client INIT\n\n")
	}
	return ctx

}
