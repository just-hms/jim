package actions

import (
	"jim/models"
	"jim/utils"
)

var Rename = &Action{
	Value: func(args []string) {
		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			return
		}

		to_rename := models.Command{}

		if err := models.DB().Where("name = ?", args[1]).First(&to_rename).Error; err == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[1])
			return
		}

		models.DB().Model(&command).Update("name", args[1])
	},
	Description:     "rename a command",
	HelpDescription: "wp",
	ArgumentsLen:    2,
}
