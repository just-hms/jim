package actions

import (
	"fmt"
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
				"Invoke-Expression",
				"'cmd /c start powershell -windowstyle hidden -c jim --bg-watch "+command.Name+" "+strings.Join(args[1:], " ")+"'",
			)
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

		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Println(err.Error())
		}

		c.Run()

	},
	Description:     "run a command and watch it",
	HelpDescription: "wp",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
}
