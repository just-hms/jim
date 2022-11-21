package actions

import (
	"jim/pkg/models"
	"jim/pkg/rainbow"
)

var Add = &Action{
	Value: func(args []string) {

		// check if a command with this name already exists
		if err := models.GetCommandByName(&models.Command{}, args[0]); err == nil {
			rainbow.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		// if not create it

		command := models.Command{
			Name: args[0],
		}

		if err := GetCommandValueFromArgs(args, &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		if err := command.Save(); err != nil {
			rainbow.Alertf("error adding the command\n")
			return
		}

	},
	Description:     "add a command",
	HelpDescription: " add a command using this syntax\n\n     jim --add command <value>\n\n If no value is specified jim will open your default editor and\n will let you insert a set of instruction in a temporary file.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}
