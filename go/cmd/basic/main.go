package main

import (
	"fmt"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
)

var (
	session_name string
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "k":
			Keys()
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
