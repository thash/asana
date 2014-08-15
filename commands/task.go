package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
)

func Task(c *cli.Context) {
	t, stories := api.Task(api.FindTaskId(c.Args()), c.Bool("verbose"))

	fmt.Printf("[ %s ] %s\n", t.Due_on, t.Name)

	if len(t.Tags) > 0 {
		fmt.Print("  Tags: ")
		for i, tag := range t.Tags {
			print(tag.Name)
			if len(t.Tags) != 1 && i != (len(t.Tags)-1) {
				print(", ")
			}
		}
		println("")
	}
	fmt.Printf("--------\n%s\n", t.Notes)

	if stories != nil {
		fmt.Println("\n----------------------------------------\n")
		for _, s := range stories {
			fmt.Printf("%s", s)
			fmt.Println("\n--------")
		}
	}
}
