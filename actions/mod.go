package actions

import (
	"jim/models"
	"jim/utils"
	"strings"
)

var Mod = &Action{
	Value: func(args []string) {
		args[1] = strings.Replace(args[1], utils.CURRENT_FOLDER_FLAG, utils.CurrentFolder(), -1)

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		command.Value = args[1]
		models.DB().Save(&command)
	},
	Description:     "modify a specified command",
	HelpDescription: "wp",
	ArgumentsLen:    2,
}

var ModById = &Action{
	Value: func(args []string) {
		args[1] = strings.Replace(args[1], utils.CURRENT_FOLDER_FLAG, utils.CurrentFolder(), -1)

		command := models.Command{}

		if err := models.DB().Where("id = ?", args[0]).First(&command).Error; err != nil {
			utils.Alertf("id not found")
			return
		}

		command.Value = args[1]
		models.DB().Save(&command)
	},
	Description:     "modify a command by id",
	HelpDescription: "wp",

	ArgumentsLen: 2,
}
