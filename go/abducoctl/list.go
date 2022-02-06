package abducoctl

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
)

func ListRemoteHostSessions(hosts map[string]RemoteHost) {
	for name, host := range hosts {
		s := time.Now()
		session_names := []string{}
		host.ParseList(SSH(host, fmt.Sprintf(`%s -l`, ABDUCO_BINARY_NAME)))
		for _, S := range host.Sessions {
			session_names = append(session_names, S.Session)
		}
		table := uitable.New()
		table.MaxColWidth = COLS
		table.Wrap = true
		table.AddRow(color.YellowString("Host:"), color.YellowString(name))
		table.AddRow(color.GreenString("Sessions:"), strings.Join(session_names, `, `))
		table.AddRow(color.GreenString("Duration:"), time.Since(s))
		fmt.Println(table)
		fmt.Println("")
	}
}
