package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "syncs"
	app.Version = Version
	app.Usage = "HTTP server app for hosting your slides."
	app.Author = "zoncoen"
	app.Email = "zoncoen@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
