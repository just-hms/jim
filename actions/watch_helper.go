package actions

import (
	"jim/models"
	"jim/utils"
	"strings"
)

var WatchHelper = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		c, err := utils.DetachedCrossCmd(
			"jim --bg-watch",
			command.Name,
			strings.Join(args[1:], " "),
		)

		if err != nil {
			utils.Alertf(err.Error())
			return
		}

		c.Run()

	},
	Description:     "run a command and watch it",
	HelpDescription: "wp",
	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},
}
