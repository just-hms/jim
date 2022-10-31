package actions

import (
	"jim/models"
	"jim/utils"
)

var Add = &Action{
	Value: func(args []string) {

		to_search := models.Command{}

		if err := models.DB().Where("name = ?", args[0]).First(&to_search).Error; err == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		command_value, err := utils.GetCommandFromArgs(args, "")

		if err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		command := models.Command{
			Name:  args[0],
			Value: command_value,
		}

		models.DB().Create(&command)

	},
	Description:     "add a command",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}
