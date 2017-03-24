package commands

import (
	"github.com/urfave/cli"
)

// Start prepares the project directory for use
var Start = cli.Command{
	Name:  "start",
	Usage: "Start a story",
	Action: func(c *cli.Context) error {
		return nil
	},
}
