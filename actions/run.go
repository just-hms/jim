package actions

import (
	"fmt"
	"jim/models"
	"jim/rainbow"
	"jim/utils"
	"os"
	"strings"
)

var Run = &Action{
	Value: func(args []string) {

		if len(args) != 1 && len(args) != 2 {
			rainbow.Alertf("wrong format!!!\n")
			return
		}

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		if len(args) == 2 {
			run(command, strings.Join(args[1:], " "))
			return
		}

		run(command, "")
	},
	Description:     "run a command (not required)",
	HelpDescription: " Run a command using this syntax\n\n     jim <--run> command\n\n Will run the specified command in your default shell.\n --run can be omitted.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}

func run(command models.Command, args string) {

	c, err := utils.CrossCmd(
		command.Value,
		args,
	)

	if err != nil {
		rainbow.Alertf("%s\n", err.Error())
		return
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
