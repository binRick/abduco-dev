package abducoctl

import (
	"fmt"
)

func ListHosts(hosts map[string]RemoteHost) {
	for name, host := range hosts {
		host.ParseList(SSH(host, fmt.Sprintf(`%s -l`, ABDUCO_BINARY_NAME)))
		l := fmt.Sprintf(`%s- %d Sessions`, name, len(host.Sessions))
		fmt.Println(l)
	}
}
