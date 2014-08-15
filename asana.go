package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.0.1"
	app.Usage = "asana cui client"

	app.Commands = []cli.Command{
		{
			Name:      "config",
			ShortName: "c",
			Usage:     "Asana configuration. Your settings will be saved in ~/.asana.yml",
			Action: func(c *cli.Context) {
				commands.Config(c)
			},
		},
		{
			Name:      "workspaces",
			ShortName: "w",
			Usage:     "get workspaces",
			Action: func(c *cli.Context) {
				commands.Workspaces(c)
			},
		},
		{
			Name:      "tasks",
			ShortName: "ts",
			Usage:     "get tasks",
			Action: func(c *cli.Context) {
				commands.Tasks(c)
			},
		},
		{
			Name:      "task",
			ShortName: "t",
			Usage:     "get a task",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "verbose, v", Usage: "verbose output"},
			},
			Action: func(c *cli.Context) {
				commands.Task(c)
			},
		},
		{
			Name:      "comment",
			ShortName: "cm",
			Usage:     "Post comment",
			Action: func(c *cli.Context) {
				commands.Comment(c)
			},
		},
	}

	app.Run(os.Args)
}
