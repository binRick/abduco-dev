package main

import (
	"fmt"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	"github.com/k0kubun/pp"
)

func main() {
	if false {
		fmt.Fprintf(os.Stdout, "%s\n", pp.Sprintf(`%s`, abducoctl.PIDs()))
	}
	fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("%s", abducoctl.JSON()))
}
