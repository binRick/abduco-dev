package abducoctl

import (
	"context"
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
		cancel()
	}()
	return ctx

}
