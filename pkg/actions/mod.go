package actions

import (
	"jim/pkg/models"
	"jim/pkg/rainbow"
)

var Mod = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		// modify the command
		if err := GetCommandValueFromArgs(args, &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		if err := command.Save(); err != nil {
			rainbow.Alertf("error modifying the command\n")
			return
		}
	},
	Description:     "modify a specified command",
	HelpDescription: " Modify a command using this syntax\n\n     jim --mod command\n\n Will open the command in your default editor and will let\n you modify it.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}
