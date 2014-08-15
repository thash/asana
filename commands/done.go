package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
)

func Done(c *cli.Context) {
	taskId := c.Args().First()
	task := api.Update(taskId, "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
