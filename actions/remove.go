package actions

import (
	"jim/models"
	"jim/utils"
)

var Remove = &Action{
	Value: func(args []string) {

		for _, arg := range args {

			command := models.Command{}

			if err := FindCommandByName(arg, &command); err != nil {
				return
			}

			models.DB().Unscoped().Delete(&command)
		}
	},
	Description:     "remove one or more command",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},
}

var RemoveById = &Action{
	Value: func(args []string) {

		for _, arg := range args {

			command := models.Command{}

			if err := models.DB().Where("id = ?", arg).First(&command).Error; err != nil {
				utils.Alertf("id not found")
				return
			}

			models.DB().Unscoped().Delete(&command)
		}

	},
	Description:     "remove one or more command by id",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},
}
