package main

import (
	"fmt"
	"os"
	"runtime"

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
			os := runtime.GOOS
			switch os {
			//			case "darwin":
			case "linux":
				if false {
					abducoctl.Ps()
				}
			default:
				if false {
					fmt.Printf("%s.\n", os)
				}
			}
		case "connect":
			if len(os.Args) > 2 {
				abducoctl.Connect(ctx, os.Args[2])
			}
		}
	}
}
