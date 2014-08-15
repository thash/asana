package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
)

func DueOn(c *cli.Context) {
	taskId := c.Args().First()
	task := api.Update(taskId, "due_on", toDate(c.Args()[1]))
	fmt.Println("set due on [ " + task.Due_on + " ] :" + task.Name)
}

// TODO: flexible due date setting.
func toDate(str string) string {
	return str
}
