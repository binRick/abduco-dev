package main

import (
	"fmt"
	"os"
	"time"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	"github.com/k0kubun/pp"
)

var (
	session_name string
	hosts        = map[string]abducoctl.RemoteHost{
		`mac`: abducoctl.RemoteHost{
			User:    `rick`,
			Host:    `127.0.0.1`,
			Name:    `mac`,
			Port:    22,
			Timeout: (time.Millisecond * 1000),
			OS:      `darwin`,
		},
		`al1`: abducoctl.RemoteHost{
			Name:    `al1`,
			User:    `root`,
			Host:    `127.0.0.1`,
			Port:    45888,
			Timeout: (time.Millisecond * 1000),
			OS:      `linux`,
		},
		`f36`: abducoctl.RemoteHost{
			Name:    `f36`,
			User:    `root`,
			Host:    `127.0.0.1`,
			Port:    49117,
			Timeout: (time.Millisecond * 1000),
			OS:      `linux`,
		},
	}
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "normalize":
			if len(os.Args) > 2 {
				host := hosts[os.Args[2]]
				abducoctl.NormalizeRemoteHost(host)
			}
		case "k":
			Keys()
		case "ssh":
			if len(os.Args) > 2 {
				host := hosts[os.Args[2]]
				stdout := abducoctl.SSH(host, fmt.Sprintf(`%s -l`, abducoctl.DST_SB_PATH))
				host.ParseList(stdout)
				pp.Println(host)
			}
		case "remote":
			if len(os.Args) > 2 {
				host := hosts[os.Args[2]]
				abducoctl.ListRemoteHostSessions(host)
			}
		case "buffer":
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
