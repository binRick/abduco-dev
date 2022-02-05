package abducoctl

import (
	"fmt"
	"log"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
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
			return fmt.Sprintf(`Session: %s (%d)
Started: %s
Buffer: 
%s
`,
				sessions[i].Session,
				sessions[i].PID,
				sessions[i].Started,
				`xxxxxxxxxxxxxxxxxxxxxxxxxxxxx`,
			)
		}))
	if err != nil {
		log.Fatal(err)
	}
	if Exists(sessions[idx[0]].Session) {
		Connect(ctx, sessions[idx[0]].Session)
	}
}
