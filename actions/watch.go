package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var Watch = &Action{
	Value: func(args []string) {
		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		if len(args) == 2 {
			watch(command, strings.Join(args[1:], " "))
			return
		}

		watch(command, "")

	},
	Description:     "run a command and watch it",
	HelpDescription: "wp",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
}

func watch(command models.Command, args string) {

	models.DB().Save(&command)

	var c *exec.Cmd

	if runtime.GOOS == "windows" {

		c = exec.Command("powershell", "-c", "Measure-Command { "+command.Value+"}", args)

	} else {

		shell, err := os.LookupEnv("SHELL")

		if !err {
			utils.Alertf("no shell found!!!")
			return
		}

		c = exec.Command("time", shell, "-c", command.Value, args)
	}

	c.Stdin = os.Stdin
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr

	start := time.Now()

	if err := c.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	elapsed := time.Since(start)

	session := models.Session{
		Elapsed:   elapsed,
		CommandID: command.ID,
	}

	models.DB().Create(&session)
}
