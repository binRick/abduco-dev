package abducoctl

import (
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

var COLS uint
var ROWS uint

func init() {

	x, y, err := terminal.Dimensions()
	if err != nil {
		panic(err)
	}
	COLS = y
	ROWS = x
}
