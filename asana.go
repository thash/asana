package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.1.1"
	app.Usage = "asana cui client"

	app.Commands = defs()
	app.Run(os.Args)
}

func defs() []cli.Command {
	return []cli.Command{
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
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
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
		{
			Name:      "done",
			Usage:     "Complete task",
			Action: func(c *cli.Context) {
				commands.Done(c)
			},
		},
		{
			Name:  "due",
			Usage: "set due date",
			Action: func(c *cli.Context) {
				commands.DueOn(c)
			},
		},
	}
}
