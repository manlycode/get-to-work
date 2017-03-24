package commands

import (
	"github.com/urfave/cli"
)

// Stop prepares the project directory for use
var Stop = cli.Command{
	Name:  "stop",
	Usage: "Stop working on a story",
	Action: func(c *cli.Context) error {
		return nil
	},
}
