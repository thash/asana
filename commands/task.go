package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
)

func Task(c *cli.Context) {
	t, stories := api.Task(api.FindTaskId(c.Args().First(), true), c.Bool("verbose"))

	fmt.Printf("[ %s ] %s\n", t.Due_on, t.Name)

	showTags(t.Tags)
	showCustomFields(t.CustomFields)

	fmt.Printf("\n%s\n", t.Notes)

	if stories != nil {
		fmt.Printf("\n----------------------------------------\n")
		for _, s := range stories {
			fmt.Printf("%s\n", s)
		}
	}
}

func showTags(tags []api.Base) {
	if len(tags) > 0 {
		fmt.Print("  Tags: ")
		for i, tag := range tags {
			print(tag.Name)
			if len(tags) != 1 && i != (len(tags)-1) {
				print(", ")
			}
		}
		println("")
	}
}

func showCustomFields(fields []api.CustomField_t) {
	if len(fields) > 0 {
		fmt.Println("\n  Custom Fields:")
		for _, field := range fields {
			if field.DisplayValue != "" {
				fmt.Printf("    %s: %s\n", field.Name, field.DisplayValue)
			}
		}
	}
}
