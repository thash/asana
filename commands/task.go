package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"../api"
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
		for i, s := range stories {
			if s.Type == "comment" {
				if i != 0 {
					fmt.Println("--------")
				}
				fmt.Printf("%s\nby %s (%s)", s.Text, s.Created_by.Name, s.Created_at)
				if i != 0 {
					fmt.Println("\n--------")
				}
			} else {
				fmt.Printf("%s (%s)\n", s.Text, s.Created_at)
			}
		}
	}
}
