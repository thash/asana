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
	if err != nil || config.Load().Api_key == "" {
		println("visit: http://app.asana.com/-/account_api")
		print("paste your api_key: ")
		var s string
		fmt.Scanf("%s", &s)

		f, _ := os.Create(utils.Home() + "/.asana.yml")
		defer f.Close()
		f.WriteString("api_key: " + s + "\n")
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
	apiKey := config.Load().Api_key
	f, _ := os.Create(utils.Home() + "/.asana.yml")
	f.WriteString("api_key: " + apiKey + "\n")
	f.WriteString("workspace: " + strconv.Itoa(ws[index].Id) + "\n")
}
