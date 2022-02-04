package main

import (
	"flag"
	"fmt"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	"github.com/k0kubun/pp"
)

var (
	session_name string
)

func main() {
	flag.StringVar(&session_name, "name", "", "session name")
	flag.Parse()
	if false {
		fmt.Fprintf(os.Stdout, "%s\n", pp.Sprintf(`%s`, abducoctl.PIDs()))
		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("%s", abducoctl.JSON()))
		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("%s", abducoctl.Path()))
	}
	fmt.Fprintf(os.Stdout, "%s\n", pp.Sprintf(`%s`, abducoctl.Names()))
	abducoctl.Connect(SessionNameString())

}
