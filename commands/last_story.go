package commands

import (
	"github.com/urfave/cli"
)

// LastStory prepares the project directory for use
var LastStory = cli.Command{
	Name:  "last_story",
	Usage: "Show the last story you worked on",
	Action: func(c *cli.Context) error {
		return nil
	},
}
