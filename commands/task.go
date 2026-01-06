package commands

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
)

func Task(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	t, stories := api.Task(taskId, c.Bool("verbose"))
	attachments := api.Attachments(taskId)

	if c.Bool("json") {
		output := map[string]interface{}{
			"task": t,
		}
		if stories != nil {
			output["stories"] = stories
		}
		if len(attachments) > 0 {
			output["attachments"] = attachments
		}
		jsonData, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling JSON: %v\n", err)
			return
		}
		fmt.Println(string(jsonData))
		return
	}

	fmt.Printf("[ %s ] %s\n", t.Due_on, t.Name)

	showTags(t.Tags)
	showCustomFields(t.CustomFields)
	showAttachments(attachments)

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

func showAttachments(attachments []api.Attachment_t) {
	if len(attachments) > 0 {
		fmt.Printf("\n  Attachments (%d):\n", len(attachments))
		for i, att := range attachments {
			fmt.Printf("    [%d] %s (GID: %s)\n", i, att.Name, att.Gid)
		}
	}
}
