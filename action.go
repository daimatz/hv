package main

import (
	"github.com/codegangsta/cli"
)

func Action(c *cli.Context) {
	if c.GlobalBool("help") {
		cli.ShowAppHelp(c)
		return
	}

	switch len(c.Args()) {
	case 0:
		cli.ShowAppHelp(c)
	case 1:
		fromStdin(c)
	case 2:
		fromFile(c)
	default:
		cli.ShowAppHelp(c)
	}
}

func fromStdin(c *cli.Context) {

}

func fromFile(c *cli.Context) {

}
