package actions

import (
	"jim/pkg/models"
	"jim/pkg/rainbow"

	"strings"
	"time"
)

var Watch = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		params := ""

		if len(args) == 2 {
			params = strings.Join(args[1:], " ")
		}

		// continue in the BackgroundSubAction
		ContinueInBackground(command, params)

	},
	Description:     "run a command in background and time it",
	HelpDescription: " Run a command in the background and time it using this syntax\n\n     jim --watch command\n\n Will launch the command in background and save its time of execution.\n The time that the command took to execute will be visible using the --show utility.\n User input and output don't work using --watch",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},

	BackgroundSubAction: func(args []string) {

		command, params, err := TakeUp(args)

		if err != nil {
			return
		}

		WatchCommand(command, params)

	},
}

func WatchCommand(command models.Command, args string) {

	// set up the session

	session := models.Session{
		Start:   time.Now(),
		Command: command.Name,
	}

	RunCommand(command, args)

	// set the difference between the end and the start time

	session.Elapsed = time.Since(session.Start)

	// save it

	if err := session.Save(); err != nil {
		rainbow.Alertf("error adding the session\n")
		return
	}
}
