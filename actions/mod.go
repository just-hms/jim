package actions

import (
	"jim/models"
	"jim/utils"
)

var Mod = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf(err.Error())
			return
		}

		command_value, err := utils.GetCommandFromArgs(args, command.Value)

		if err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		command.Value = command_value
		models.DB().Save(&command)
	},
	Description:     "modify a specified command",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}

var ModById = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := models.DB().Where("id = ?", args[0]).First(&command).Error; err != nil {
			utils.Alertf("specified id not found")
			return
		}

		command_value, err := utils.GetCommandFromArgs(args, command.Value)

		if err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		command.Value = command_value
		models.DB().Save(&command)
	},
	Description:     "modify a command by id",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}
