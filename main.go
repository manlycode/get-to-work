package main

import (
	"github.com/manlycode/get-to-work/commands"
	"github.com/urfave/cli"
	"os"
)

const version = "0.0.2"

func main() {
	app := cli.NewApp()
	app.Name = "get-to-work"
	app.Version = version
	app.Description = "Keep track of what you're doing in Harvest"

	app.Commands = []cli.Command{
		commands.Bootstrap,
		commands.Start,
		commands.Stop,
		commands.LastStory,
	}

	app.Run(os.Args)
}
