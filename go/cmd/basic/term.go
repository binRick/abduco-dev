package main

import (
	"context"
	"log"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	"github.com/pkg/term"
	"github.com/tj/go-terminput"
)

func Term(ctx context.Context) {
	//go func() {
	for {
		if len(abducoctl.Names()) < 1 {
			abducoctl.Connect(ctx, abducoctl.NewNameString())
		} else {
			abducoctl.Prompt()
			func() {
				t, err := term.Open("/dev/tty")
				if err != nil {
					log.Fatalf("error: %s\n", err)
				}
				t.SetRaw()
				defer t.Restore()

				e, err := terminput.Read(t)
				if err != nil {
					panic(err)
				}
				//fmt.Fprintf(os.Stderr, "\n\n%s\n\n", e.Key())
				is_ctrl_c := e.Key() == 3
				if is_ctrl_c {
					os.Exit(0)
				}
				is_ctrl_n := e.Key() == 14
				if is_ctrl_n {
					abducoctl.Connect(ctx, abducoctl.NewNameString())
				}
			}()
		}
	}
	//	}()
}
