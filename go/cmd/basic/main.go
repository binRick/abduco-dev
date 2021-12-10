package main

import (
	"encoding/json"
	"fmt"
	"os"

	abducoctl "github.com/binRick/abduco-dev/go/abducoctl"
)

func main() {
	l, _ := abducoctl.List()
	dat, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", string(dat))
}
