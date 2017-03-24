package commands

import (
	"fmt"
	"github.com/manlycode/get-to-work/services/harvest"
	"github.com/urfave/cli"
	"log"
)

// Bootstrap prepares the project directory for use
var Bootstrap = cli.Command{
	Name:  "bootstrap",
	Usage: "Prepare the current project directory for get-to-work",
	Action: func(c *cli.Context) error {
		harvestService, err := harvest.NewService("foo", "bar@example.com", "baz")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(harvestService.User.ID)

		return nil
	},
}
