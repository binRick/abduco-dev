package main

import (
	//abduco "local.dev/abduco"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	abduco "local.dev/abduco"
)

func main() {
	fmt.Println("vim-go")
	l, _ := abduco.List()
	fmt.Fprintf(os.Stderr, "%s\n", pp.Sprintf("%s", l))
}
