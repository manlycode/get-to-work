package commands

import (
	"github.com/urfave/cli"
)

// Bootstrap prepares the project directory for use
var Bootstrap = cli.Command{
	Name:  "bootstrap",
	Usage: "Prepare the current project directory for get-to-work",
	Action: func(c *cli.Context) error {
		return nil
	},
}
