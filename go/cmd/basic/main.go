package main

import (
	"fmt"
	"os"
	"time"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
)

var (
	session_name string
	f36          = abducoctl.RemoteHost{
		User:    `root`,
		Host:    `127.0.0.1`,
		Port:    49117,
		Timeout: (time.Millisecond * 1000),
	}
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "k":
			Keys()
		case "ssh":
			stdout := abducoctl.SSH(f36, `abduco-sb -l`)
			fmt.Println(stdout)
		case "dev":
			Dev()
		case "b":
			if len(os.Args) > 2 {
				lines := abducoctl.Buffer(os.Args[2])
				for _, l := range lines {
					fmt.Fprintf(os.Stdout, "%s\n", l)
				}
			}
		case "find":
			abducoctl.Finder()
		case "list":
			abducoctl.List()
		case "select":
			abducoctl.Select()
		case "ps":
			abducoctl.Ps()
		case "connect":
			if len(os.Args) > 2 {
				abducoctl.Connect(ctx, os.Args[2])
			}
		}
	}
}
