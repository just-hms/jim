package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"strings"
	"time"
)

var Watch = &Action{
	Value: func(args []string) {
		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf(err.Error())
			return
		}

		if len(args) == 2 {
			watch(command, strings.Join(args[1:], " "))
			return
		}

		watch(command, "")

	},
	Description:     "run a command in background and time it (user input and output don't work)",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},

	BackGround: true,
}

func watch(command models.Command, args string) {

	c, err := utils.CrossCmd(
		command.Value,
		args,
	)

	if err != nil {
		utils.Alertf(err.Error())
		return
	}

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
