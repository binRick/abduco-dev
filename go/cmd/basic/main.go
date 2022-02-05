package main

import (
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
)

var (
	session_name string
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "f":
			abducoctl.Finder()
		//		case "b":
		//			abducoctl.Buffer()
		case "list":
			abducoctl.List()
		case "select":
			abducoctl.Prompt()
		case "ps":
			abducoctl.Ps()
		case "connect":
			if len(os.Args) > 2 {
				abducoctl.Connect(ctx, os.Args[2])
			}
		}
	}
}
