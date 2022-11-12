package actions

import (
	"jim/models"
	"jim/utils"
)

var Mod = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		if err := utils.GetCommandValueFromArgs(args, &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		if err := command.Save(); err != nil {
			utils.Alertf("error modifying the command\n")
			return
		}
	},
	Description:     "modify a specified command",
	HelpDescription: " Modify a command using this syntax\n\n     jim --mod command\n\n Will open the command in your default editor and will let\n you modify it.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}
