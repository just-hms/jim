package actions

import (
	"jim/models"
	"jim/utils"
	"strconv"

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
}

var BgWatch = &Action{
	Value: func(args []string) {

		command_id, _ := strconv.ParseUint(args[0], 10, 32)
		command_value := args[1]
		args_string := strings.Join(args[2:], " ")

		watch(
			command_value,
			uint(command_id),
			args_string,
		)

	},
	Description:     "utility action, don't call this from the command line",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 2
	},
}

func watch(command_value string, command_id uint, args string) {

	c, err := utils.CrossCmd(command_value, args)

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
		CommandID: command_id,
	}

	models.DB().Create(&session)
}
