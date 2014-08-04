package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
)

type Task struct {
	Id              int
	Created_at      string
	Modified_at     string
	Name            string
	Notes           string
	Assignee        Plain
	Completed       bool
	Assignee_status string
	Completed_at    string
	Due_on          string
	Tags            []Plain
	Workspace       Plain
	Parent          string
	Projects        []Plain
	Folloers        []Plain
}

type Story struct {
	Id         int
	Text       string
	Type       string
	Created_at string
	Created_by Plain
}

type Me struct {
	Id         int
	Name       string
	Email      string
	Workspaces []Plain
	Photo      Images
}

type Images struct {
	Image_128x128 string
	Image_60x60   string
	Image_36x36   string
	Image_27x27   string
	Image_21x21   string
}

type Plain struct {
	Id   int
	Name string
}

// http://jordanorelli.com/post/42369331748/function-types-in-go-golang
func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.0.1"
	app.Usage = "asana cui client"

	app.Commands = []cli.Command{
		{
			Name:      "workspaces",
			ShortName: "w",
			Usage:     "get workspaces",
			Action: func(c *cli.Context) {
				for _, w := range me().Workspaces {
					fmt.Printf("%16d %s\n", w.Id, w.Name)
				}
			},
		},
		{
			Name:      "tasks",
			ShortName: "t",
			Usage:     "get tasks",
			Action: func(c *cli.Context) {
				params := url.Values{}
				params.Add("workspace", strconv.Itoa(load_config().Workspace))
				params.Add("assignee", "me")
				params.Add("opt_fields", "name,completed,due_on")
				for _, t := range tasks(params) {
					if t.Completed {
						continue
					}
					fmt.Printf("%16d [ %10s ] %s\n", t.Id, t.Due_on, t.Name)
				}
			},
		},
		{
			Name:      "show",
			ShortName: "s",
			Usage:     "get a task",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "verbose, v", Usage: "verbose output"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fmt.Printf("%s", "usage: $ asana show 123\n")
					return
				}

				task, stories := task(c.Args().First(), c.Bool("verbose"))

				fmt.Printf("[ %s ] %s\n", task.Due_on, task.Name)

				if len(task.Tags) > 0 {
					fmt.Print("  Tags: ")
					for i, tag := range task.Tags {
						print(tag.Name)
						if len(task.Tags) != 1 && i != (len(task.Tags)-1) {
							print(", ")
						}
					}
					println("")
				}
				fmt.Printf("--------\n%s\n", task.Notes)

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
			},
		},
	}

	app.Run(os.Args)
}

func get(path string, params url.Values) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", build_get_url("https://app.asana.com", path, params), nil)
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36"
	req.Header.Set("User-Agent", ua)
	req.SetBasicAuth(load_config().Api_key, "")
	resp, err := client.Do(req)
	fatal(err)

	contents, err2 := ioutil.ReadAll(resp.Body)
	fatal(err2)

	return contents
}

func build_get_url(host string, path string, params url.Values) string {
	if params == nil || params.Encode() == "" {
		return host + path
	} else {
		return host + path + "?" + params.Encode()
	}
}

type Conf struct {
	Api_key   string
	Workspace int
}

// TODO: search ~/.asana.yml
func load_config() Conf {
	dat, err := ioutil.ReadFile(".asana.yml")
	fatal(err)
	conf := Conf{}
	err2 := yaml.Unmarshal(dat, &conf)
	fatal(err2)
	return conf
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func me() Me {
	var me map[string]Me
	err := json.Unmarshal(get("/api/1.0/users/me", nil), &me)
	fatal(err)
	return me["data"]
}

type ByDue []Task

func (a ByDue) Len() int           { return len(a) }
func (a ByDue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDue) Less(i, j int) bool { return a[i].Due_on < a[j].Due_on }

func tasks(params url.Values) []Task {
	var tasks map[string][]Task
	err := json.Unmarshal(get("/api/1.0/tasks", params), &tasks)
	fatal(err)
	var tasks_without_due, tasks_with_due []Task
	for _, t := range tasks["data"] {
		if t.Due_on == "" {
			tasks_without_due = append(tasks_without_due, t)
		} else {
			tasks_with_due = append(tasks_with_due, t)
		}
	}
	sort.Sort(ByDue(tasks_with_due))
	return append(tasks_with_due, tasks_without_due...)
}

func task(taskId string, verbose bool) (Task, []Story) {
	var t map[string]Task
	err := json.Unmarshal(get("/api/1.0/tasks/"+taskId, nil), &t)
	fatal(err)
	if verbose {
		var stories map[string][]Story
		err2 := json.Unmarshal(get("/api/1.0/tasks/"+taskId+"/stories", nil), &stories)
		fatal(err2)
		return t["data"], stories["data"]
	} else {

		return t["data"], nil
	}
}
