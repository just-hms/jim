package actions

import (
	"jim/models"
	"jim/utils"
)

var Remove = &Action{
	Value: func(args []string) {
		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		models.DB().Unscoped().Delete(&command)
	},
	Description:     "remove a command",
	HelpDescription: "wp",

	ArgumentsLen: 1,
}

var RemoveById = &Action{
	Value: func(args []string) {
		command := models.Command{}

		if err := models.DB().Where("id = ?", args[0]).First(&command).Error; err != nil {
			utils.Alertf("id not found")
			return
		}

		models.DB().Unscoped().Delete(&command)
	},
	Description:     "remove a command by id",
	HelpDescription: "wp",

	ArgumentsLen: 1,
}
