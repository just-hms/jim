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
	Description:     "run a command in background and time it",
	HelpDescription: " Run a command in the background and time it using this syntax\n\n     jim --watch command\n\n Will launch the command in background and save its time of execution.\n The time that the command took to execute will be visible using the --show utility.\n User input and output don't work using --watch",

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

		session := models.Session{
			Start:   time.Now(),
			Command: command.Name,
		}

		if err := c.Run(); err != nil {
			return
		}

		session.Elapsed = time.Since(session.Start)

		if err := session.Save(); err != nil {
			utils.Alertf("error adding the session\n")
			return
		}
	},
}
