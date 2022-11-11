package actions

import (
	"jim/models"
	"jim/utils"
)

var Rename = &Action{
	Value: func(args []string) {

		command := models.Command{}

		// find the command
		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		// check if a command with the new_name exists
		if err := models.GetCommandByName(&models.Command{}, args[1]); err == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		if err := command.Rename(args[1]); err != nil {
			utils.Alertf("error renaming the command\n")
			return
		}

	},
	Description:     "rename a command",
	HelpDescription: " Rename a command using this syntax\n\n     jim --rn command new_name\n\n Will rename the specified command with the provided `new_name`.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 2
	},
}
