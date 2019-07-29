package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strconv"

	"github.com/memerelics/asana/config"
	"github.com/memerelics/asana/utils"
)

type Task_t struct {
	Id              int
	Created_at      string
	Modified_at     string
	Name            string
	Notes           string
	Assignee        Base
	Completed       bool
	Assignee_status string
	Completed_at    string
	Due_on          string
	Tags            []Base
	Workspace       Base
	Parent          Base
	Projects        []Base
	Folloers        []Base
}

type Story_t struct {
	Id         int
	Text       string
	Type       string
	Created_at string
	Created_by Base
}

type ByDue []Task_t

func (a ByDue) Len() int           { return len(a) }
func (a ByDue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDue) Less(i, j int) bool { return a[i].Due_on < a[j].Due_on }

func Tasks(params url.Values, withCompleted bool) []Task_t {
	params.Add("workspace", strconv.Itoa(config.Load().Workspace))
	params.Add("assignee", "me")
	params.Add("opt_fields", "name,completed,due_on")
	var tasks map[string][]Task_t
	err := json.Unmarshal(Get("/api/1.0/tasks", params), &tasks)
	utils.Check(err)
	var tasks_without_due, tasks_with_due []Task_t
	for _, t := range tasks["data"] {
		if !withCompleted && t.Completed {
			continue
		}
		if t.Due_on == "" {
			tasks_without_due = append(tasks_without_due, t)
		} else {
			tasks_with_due = append(tasks_with_due, t)
		}
	}
	sort.Sort(ByDue(tasks_with_due))
	return append(tasks_with_due, tasks_without_due...)
}

func Task(taskId string, verbose bool) (Task_t, []Story_t) {
	var (
		err     error
		t       map[string]Task_t
		ss      map[string][]Story_t
		stories []Story_t
	)
	task_chan, stories_chan := make(chan []byte), make(chan []byte)
	go func() {
		task_chan <- Get("/api/1.0/tasks/"+taskId, nil)
	}()

	stories = nil
	if verbose {
		go func() {
			stories_chan <- Get("/api/1.0/tasks/"+taskId+"/stories", nil)
		}()
		err = json.Unmarshal(<-stories_chan, &ss)
		utils.Check(err)
		stories = ss["data"]
	}

	err = json.Unmarshal(<-task_chan, &t)
	utils.Check(err)
	return t["data"], stories
}

func FindTaskId(index string, autoFirst bool) string {
	if index == "" {
		if autoFirst == false {
			log.Fatal("fatal: Task index is required.")
		} else {
			index = "0"
		}
	}

	var id string
	txt, err := ioutil.ReadFile(utils.CacheFile())

	if err != nil { // cache file not exist
		ind, parseErr := strconv.Atoi(index)
		utils.Check(parseErr)
		task := Tasks(url.Values{}, false)[ind]
		id = strconv.Itoa(task.Id)
	} else {
		lines := regexp.MustCompile("\n").Split(string(txt), -1)
		for i, line := range lines {
			if index == strconv.Itoa(i) {
				line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
				id = regexp.MustCompile("^[0-9]*").FindString(line)
			}
		}
	}
	return id
}

func (s Story_t) String() string {
	if s.Type == "comment" {
		return fmt.Sprintf("> %s\nby %s (%s)", s.Text, s.Created_by.Name, s.Created_at)
	} else {
		return fmt.Sprintf("* %s (%s)", s.Text, s.Created_at)
	}
}

type Commented_t struct {
	Text string `json:"text"` // Define only required field.
}

func CommentTo(taskId string, comment string) string {

	respBody := Post("/tasks/"+taskId+"/stories", `{"data":{"text":"`+comment+`"}}`)

	var output map[string]Commented_t
	err := json.Unmarshal(respBody, &output)
	utils.Check(err)

	return output["data"].Text
}

func Update(taskId string, key string, value string) Task_t {
	respBody := Put("/tasks/"+taskId, `{"data":{"`+key+`":"`+value+`"}}`)

	var output map[string]Task_t
	err := json.Unmarshal(respBody, &output)
	utils.Check(err)

	return output["data"]
}
