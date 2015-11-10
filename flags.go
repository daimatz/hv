package main

import (
	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	cli.BoolFlag{
		Name:  "help, h",
		Usage: "show help",
	},
	cli.StringFlag{
		Name:  "separator, s",
		Value: ",",
		Usage: "set the separator to output",
	},
	cli.StringSliceFlag{
		Name:  "only, o",
		Usage: "specify fields to show. if given, output streamingly.",
	},
}
