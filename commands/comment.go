package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli"

	"github.com/memerelics/asana/api"
	"github.com/memerelics/asana/utils"
)

func Comment(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), false)
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

	postComment := trim(string(txt))
	if postComment != "" {
		commented := api.CommentTo(taskId, postComment)
		fmt.Println("Commented on Task: \"" + task.Name + "\"\n")
		fmt.Println(commented)
	} else {
		fmt.Println("Aborting comment due to empty content.")
	}
}

func template(f *os.File, task api.Task_t, stories []api.Story_t) error {
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
