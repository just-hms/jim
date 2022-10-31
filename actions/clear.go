package actions

import (
	"jim/models"
	"jim/utils"
)

var Clear = &Action{
	Value: func(args []string) {

		utils.Alertf("clear all commands is not reversible, are you sure? Type y or N\n")

		if utils.ReadChar() != 'y' {
			return
		}

		models.DB().Unscoped().Where("1=1").Delete(&models.Command{})
	},
	Description:     "clear all commands",
	HelpDescription: "wp",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
}
