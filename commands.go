package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandInit,
	commandUp,
}

var commandInit = cli.Command{
	Name:  "init",
	Usage: "",
	Description: `
`,
	Action: doInit,
}

var commandUp = cli.Command{
	Name:  "up",
	Usage: "",
	Description: `
`,
	Action: doUp,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doInit(c *cli.Context) {
}

func doUp(c *cli.Context) {
}
