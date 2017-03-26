package commands

import (
	"fmt"
	"github.com/manlycode/go-to-work/config"
	"github.com/manlycode/go-to-work/service"
	"github.com/urfave/cli"
	"log"
)

// Bootstrap prepares the project directory for use
var Bootstrap = cli.Command{
	Name:  "bootstrap",
	Usage: "Prepare the current project directory for go-to-work",
	Action: func(c *cli.Context) error {
		harvestService := service.NewHarvestService()
		err := harvestService.SignIn(
			"foo_bar",
			"user@example.com",
			"password",
		)

		cfg := config.Config{
			Services: map[string]*service.Service{
				"harvest": &harvestService.Service,
			},
		}

		cfg.Save()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(harvestService.User.ID)

		return nil
	},
}
