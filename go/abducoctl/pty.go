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
	"github.com/k0kubun/pp"
	"golang.org/x/term"
)

var (
	CMD_KEY             = `ABDUCO_CMD`
	SCROLL_BUFFER_LINES = 50
	ABDUCO_SESSION      = `A100`
	CMD                 = `/usr/bin/env bash -i`
)

func Connect(session_name string) error {
	c := exec.Command(Path(), `-L`, fmt.Sprintf(`%d`, SCROLL_BUFFER_LINES), `-A`, session_name)
	c.Env = os.Environ()
	c.Env = append(c.Env, fmt.Sprintf("%s=%s", CMD_KEY, CMD))
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	defer func() {
		_ = ptmx.Close()
	}()
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
	defer func() {
		signal.Stop(ch)
		close(ch)
	}()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = term.Restore(int(os.Stdin.Fd()), oldState)
		pp.Println(`names:`, Names())
		pp.Println(`pids:`, PIDs())
		//list, _ := List()
		//		pp.Println(`pids:`, list)
		if Exists(session_name) {
			fmt.Fprintf(os.Stderr, "You can reconnect to this session with %s\n", session_name)
		}
	}()
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}