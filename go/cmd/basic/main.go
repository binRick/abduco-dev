package main

import (
	"fmt"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
	"github.com/k0kubun/pp"
)

func main() {
	fmt.Println("vim-go")
	l, _ := abducoctl.List()
	fmt.Fprintf(os.Stderr, "%s\n", pp.Sprintf("%s", l))
}
