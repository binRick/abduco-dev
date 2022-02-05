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
Username: %s
Started: %s (%s ago)
Processes: %d
Threads: %d
---
%s
`,
				sessions[i].Session, sessions[i].PID,
				sessions[i].Username,
				sessions[i].Started, sessions[i].Duration,
				len(sessions[i].PIDs),
				sessions[i].Threads,
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
