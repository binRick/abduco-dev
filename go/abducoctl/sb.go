package abducoctl

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/leaanthony/go-ansi-parser"
)

var (
	GET_BUFFER_SESSION_SCRIPT = `./../../abducoctl/get_session_buffer.sh`
)

func PlainBuffer(n string) []string {
	s := []string{}
	for _, l := range Buffer(n) {
		clean, e := ansi.Cleanse(l)
		if e == nil {
			s = append(s, clean)
		} else {
			s = append(s, l)
		}
	}
	return s
}

func Buffer(name string) []string {
	cmd := exec.Command(GET_BUFFER_SESSION_SCRIPT, name)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return strings.Split(string(stdout.Bytes()), "\n")
}
