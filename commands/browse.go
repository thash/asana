package commands

import (
	"strconv"
	"os/exec"

    "github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
	"github.com/memerelics/asana/config"
	"github.com/memerelics/asana/utils"
)

func Browse(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
