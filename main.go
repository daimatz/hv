package main

import (
	"github.com/codegangsta/cli"

	"os"
)

var Version string = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "hv"
	app.Usage = "parse & format HTTP requests"
	app.Version = Version
	app.Author = "daimatz"
	app.Email = "dai@daimatz.net"
	app.Flags = Flags
	app.HideHelp = true
	app.Action = Action
	app.Run(os.Args)
}
