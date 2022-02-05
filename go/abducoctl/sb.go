package abducoctl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Buffer(name string) []string {
	cmd := exec.Command(`./../../abducoctl/get_session_buffer.sh`, name)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if false {
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	}
	return strings.Split(outStr, "\n")
}
