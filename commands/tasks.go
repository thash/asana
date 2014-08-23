package commands

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/memerelics/asana/api"
	"github.com/memerelics/asana/utils"
)

const (
	CacheDuration = "5m"
)

func Tasks(c *cli.Context) {
	if c.Bool("no-cache") {
		fromAPI(false)
	} else {
		if utils.Older(CacheDuration, utils.CacheFile()) || c.Bool("refresh") {
			fromAPI(true)
		} else {
			txt, err := ioutil.ReadFile(utils.CacheFile())
			if err == nil {
				lines := regexp.MustCompile("\n").Split(string(txt), -1)
				for _, line := range lines {
					format(line)
				}
			} else {
				fromAPI(true)
			}
		}
	}
}

func fromAPI(saveCache bool) {
	tasks := api.Tasks(url.Values{}, false)
	if saveCache {
		cache(tasks)
	}
	for i, t := range tasks {
		fmt.Printf("%2d [ %10s ] %s\n", i, t.Due_on, t.Name)
	}
}

func cache(tasks []api.Task_t) {
	f, _ := os.Create(utils.CacheFile())
	defer f.Close()
	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(t.Id) + ":")
		f.WriteString(t.Due_on + ":")
		f.WriteString(t.Name + "\n")
	}
}

func format(line string) {
	index := regexp.MustCompile("^[0-9]*").FindString(line)
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove task_id
	date := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}").FindString(line)
	line = regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}:").ReplaceAllString(line, "") // remove date
	fmt.Printf("%2s [ %10s ] %s\n", index, date, line)
}
