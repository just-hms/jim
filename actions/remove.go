package actions

import (
	"jim/models"
	"jim/utils"
)

var Remove = &Action{
	Value: func(args []string) {

		if len(args) < 1 {
			utils.Alertf("wrong format!!!")
		}

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

	ArgumentsLen: -1,
}

var RemoveById = &Action{
	Value: func(args []string) {

		if len(args) < 1 {
			utils.Alertf("wrong format!!!")
		}

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

	ArgumentsLen: -1,
}
