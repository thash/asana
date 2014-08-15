package commands

import (
	"fmt"
	"net/url"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
)

func Tasks(c *cli.Context) {
	for _, t := range api.Tasks(url.Values{}, false) {
		fmt.Printf("%16d [ %10s ] %s\n", t.Id, t.Due_on, t.Name)
	}
}
