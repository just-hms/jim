package actions

import (
	"jim/constants"
	"jim/models"
	"jim/rainbow"
	"jim/utils"
)

var Clear = &Action{
	Value: func(args []string) {

		// if force is set launch clear without asking
		if len(args) == 1 && args[0] == constants.ACTION_PREFIX+"force" {
			models.Clear()
			return
		}

		rainbow.Alertf("clear all commands is not reversible, are you sure? Type y or N\n")

		if utils.ReadChar() != 'y' {
			return
		}

		models.Clear()

	},
	Description:     "clear all commands",
	HelpDescription: " Clear all commands using this syntax\n\n     jim --clear\n\n Will remove all commands.",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0 || len(args) == 1 && args[0] == constants.ACTION_PREFIX+"force"
	},
}
