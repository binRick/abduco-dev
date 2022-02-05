package abducoctl

import (
	"fmt"
	"log"

	pp "github.com/k0kubun/pp"
	"github.com/leaanthony/go-ansi-parser"
	fuzzyfinder "local.dev/go-fuzzyfinder"
)

func Finder() {
	sessions, _ := List()
	idx, err := fuzzyfinder.Find(
		sessions,
		func(i int) string {
			return sessions[i].Session
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			sess := pp.Sprintf(`%s`, sessions[i])
			sess_c, err := ansi.Cleanse(sess)
			if err != nil {
				panic(err)
			}
			return fmt.Sprintf(`Session: %s (%d)
Started: %s (%s ago)
---
%s
`,
				sessions[i].Session,
				sessions[i].PID,
				sessions[i].Started, sessions[i].Duration,
				sess_c,
			)
		}))
	if err != nil {
		log.Fatal(err)
	}
	if Exists(sessions[idx].Session) {
		Connect(ctx, sessions[idx].Session)
	}
}
