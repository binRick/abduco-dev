package abducoctl

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"golang.org/x/term"
)

var (
	CMD_KEY             = `ABDUCO_CMD`
	SCROLL_BUFFER_LINES = 50
	ABDUCO_SESSION      = `A100`
)

func Connect() error {
	cmd := `/usr/bin/env bash -i`
	c := exec.Command(Path(), `-L`, fmt.Sprintf(`%d`, SCROLL_BUFFER_LINES), `-A`, ABDUCO_SESSION)
	c.Env = os.Environ()
	c.Env = append(c.Env, fmt.Sprintf("%s=%s", CMD_KEY, cmd))
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	defer func() { _ = ptmx.Close() }() // Best effort.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH
	defer func() { signal.Stop(ch); close(ch) }()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}
