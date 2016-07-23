package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "tf2"
	app.Usage = "Explore a Docker registry"

	app.Commands = []cli.Command{
		{
			Name:  "tags",
			Usage: "list tags for a repository",
			Action: func(c *cli.Context) error {
				ref := c.Args().First()
				fmt.Println("get tags for", ref)
				return nil
			},
		},
	}

	app.Run(os.Args)

}
