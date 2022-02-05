package abducoctl

import (
	"fmt"
	"log"
	"strings"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/leaanthony/go-ansi-parser"
)

func Finder() {
	sessions, _ := List()
	idx, err := fuzzyfinder.FindMulti(
		sessions,
		func(i int) string {
			return sessions[i].Session
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			buf := strings.Join(PlainBuffer(sessions[i].Session), "\n")
			text, err := ansi.Cleanse("\u001b[1;31;40mHello World\033[0m")
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf(`Session: %s (%d)
Started: %s
Buffer: 
---
%s
---
%s
`,
				sessions[i].Session,
				sessions[i].PID,
				sessions[i].Started,
				buf,
				text,
			)
		}))
	if err != nil {
		log.Fatal(err)
	}
	if Exists(sessions[idx[0]].Session) {
		Connect(ctx, sessions[idx[0]].Session)
	}
}
