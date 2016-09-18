package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
	"github.com/memerelics/asana/config"
	"github.com/memerelics/asana/utils"
)

func Config(c *cli.Context) {
	_, err := ioutil.ReadFile(utils.Home() + "/.asana.yml")
	if err != nil || config.Load().Personal_access_token == "" {
		println("visit: http://app.asana.com/-/account_api")
		println("  Settings > Apps > Manage Developer Apps > Personal Access Tokens")
		println("  + Create New Personal Access Token")
		print("\npaste your Personal Access Token: ")
		var s string
		fmt.Scanf("%s", &s)

		f, _ := os.Create(utils.Home() + "/.asana.yml")
		defer f.Close()
		f.WriteString("personal_access_token: " + s + "\n")
	}

	ws := api.Me().Workspaces
	index := 0

	if len(ws) > 1 {
		fmt.Println("\n" + strconv.Itoa(len(ws)) + " workspaces found.")
		for i, w := range ws {
			fmt.Printf("[%d] %16d %s\n", i, w.Id, w.Name)
		}
		index = utils.EndlessSelect(len(ws)-1, index)
	}
	token := config.Load().Personal_access_token
	f, _ := os.Create(utils.Home() + "/.asana.yml")
	f.WriteString("personal_access_token: " + token + "\n")
	f.WriteString("workspace: " + strconv.Itoa(ws[index].Id) + "\n")
}
