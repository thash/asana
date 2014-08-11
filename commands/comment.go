package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"

	"../api"
	"../utils"
)

func Comment(c *cli.Context) {
	taskId := api.FindTaskId(c.Args())
	task, stories := api.Task(taskId, true)

	tmpFile := os.TempDir() + "/asana_comment.txt"
	f, err := os.Create(tmpFile)
	utils.Check(err)
	defer f.Close()

	err = template(f, task, stories)
	utils.Check(err)

	cmd := exec.Command(os.Getenv("EDITOR"), tmpFile)
	cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
	err = cmd.Run()

	txt, err := ioutil.ReadFile(tmpFile)

	utils.Check(err)

	commented := api.CommentTo(taskId, trim(string(txt)))

	fmt.Println("Commented on Task: \"" + task.Name + "\"\n")
	fmt.Println(commented)
}

func template(f *os.File, task api.Task_t, stories []api.Story) error {
	var err error
	_, err = f.WriteString("\n\n\n")
	_, err = f.WriteString("# =================================== \n")
	_, err = f.WriteString("# " + task.Name + "\n#\n")
	_, err = f.WriteString(commentOut(task.Notes) + "\n#\n")
	_, err = f.WriteString("\n# ----------------------------------- \n")
	for _, s := range stories {
		_, err = f.WriteString(commentOut(fmt.Sprintf("%s", s)) + "\n")
	}
	return err
}

func commentOut(txt string) string {
	return strings.Replace("# "+txt, "\n", "\n# ", -1)
}

func trim(txt string) string {
	var result string
	result = regexp.MustCompile("#.*\n").ReplaceAllString(txt, "")    // Remove comments
	result = regexp.MustCompile("\n*$").ReplaceAllString(result, "")  // Remove blank lines
	result = regexp.MustCompile("\n").ReplaceAllString(result, "\\n") // Escape
	return result
}
