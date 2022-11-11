package actions

import (
	"jim/models"
	"jim/utils"
)

var Add = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := models.GetCommand(&command, args[0]); err == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		if err := utils.GetCommandFromUser(args, &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		if err := command.Save(); err != nil {
			utils.Alertf("error adding the command\n")
			return
		}

	},
	Description:     "add a command",
	HelpDescription: " add a command using this syntax\n\n     jim --add command <value>\n\n If no value is specified jim will open your default editor and\n will let you insert a set of instruction in a temporary file.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}
