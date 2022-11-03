package actions

import (
	"jim/models"
	"jim/utils"

	"strings"
	"time"
)

var Watch = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		params := ""

		if len(args) == 2 {
			params = strings.Join(args[1:], " ")
		}

		utils.ContinueInBackGround(command, params)

	},
	Description:     "run a command in background and time it (user input and output don't work)",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},

	BackgroundShit: func(args []string) {

		command, params, err := utils.TakeUp(args)

		if err != nil {
			return
		}

		c, err := utils.CrossCmd(command.Value, params)

		if err != nil {
			return
		}

		start := time.Now()

		if err := c.Run(); err != nil {
			return
		}

		elapsed := time.Since(start)

		session := models.Session{
			Elapsed:   elapsed,
			CommandID: command.ID,
		}

		models.DB().Create(&session)
	},
}
