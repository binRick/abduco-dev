// +build darwin
package main

import (
	"fmt"

	lp "github.com/tejasmanohar/go-libproc"
)

func init() {

	fmt.Println(lp.ListPids(1, 0, 0))
	fmt.Println(lp.ListAllPids(0))
}
