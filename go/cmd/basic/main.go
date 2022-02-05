package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	lp "github.com/tejasmanohar/go-libproc"
)

var (
	session_name string
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "select":
			abducoctl.Prompt()
		case "ps":
			os := runtime.GOOS
			switch os {
			case "darwin":
				if false {
					fmt.Println("MAC operating system")
					pids, e := lp.ListAllPids(0)
					if e != nil {
						panic(e)
					}
					for _, p := range pids {
						fmt.Println(p)
					}
					time.Sleep(10 * time.Second)
				}
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
