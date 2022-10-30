package actions

import (
	"jim/models"
	"jim/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var WatchHelper = &Action{
	Value: func(args []string) {

		if len(args) < 1 {
			utils.Alertf("wrong format!!!")
			return
		}

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		var c *exec.Cmd

		if runtime.GOOS == "windows" {

			c = exec.Command(
				"powershell",
				"-c",
				"Start-Process -NoNewWindow jim \"--bg-watch "+strings.Join(args, " ")+"\"")
		} else {

			shell, err := os.LookupEnv("SHELL")

			if !err {
				utils.Alertf("no shell found!!!")
				return
			}

			c = exec.Command(
				shell,
				"-c",
				"nohup jim --bg-watch "+strings.Join(args, " ")+" &")
		}

		c.Run()

	},
	Description:     "run a command and watch it",
	HelpDescription: "wp",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
}
