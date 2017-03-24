package main

import (
	"github.com/manlycode/get-to-work/version"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "get-to-work"
	app.Version = version.VERSION
	app.Description = "Keep track of what you're doing in Harvest"
	app.Run(os.Args)
}
