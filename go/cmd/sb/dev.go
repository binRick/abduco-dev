package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/k0kubun/pp"
	"github.com/pkg/term"
	"github.com/tj/go-terminput"
)

type Mod int
type Key rune
type KeyboardInput struct {
	k Key
	r rune
	m Mod
}

func Dev() {
	fmt.Fprintf(os.Stderr, "tcell.KeyEsc=%v\n", tcell.KeyEsc)
	fmt.Fprintf(os.Stderr, "tcell.KeyEsc=%v\n", tcell.NewEventKey(tcell.KeyEsc, '0', tcell.ModNone))

	ctrl_slash := &KeyboardInput{28, 28, 1}
	pp.Fprintf(os.Stderr, "%s\n", ctrl_slash)

}

func DevTTY() {
	t, err := term.Open("/dev/tty")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	t.SetRaw()
	defer t.Restore()

	fmt.Printf("Type something, use 'q' to exit.\r\n")

	for {
		e, err := terminput.Read(t)
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}

		if e.Key() == terminput.KeyEscape || e.Rune() == 'q' {
			break
		}

		fmt.Printf("%s — shift=%v ctrl=%v alt=%v meta=%v\r\n", e.String(), e.Shift(), e.Ctrl(), e.Alt(), e.Meta())
		pp.Fprintf(os.Stderr, "%s\n", e)
		pp.Fprintf(os.Stderr, "key:%s\n", e.Key())
		pp.Fprintf(os.Stderr, "rune:%v\n", e.Rune())
	}

}
