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
				utils.Alertf("%s\n", err.Error())
				return
			}

			if err := command.Remove(); err != nil {
				utils.Alertf("%s\n", err.Error())
			}

		}
	},
	Description:     "remove one or more specified command",
	HelpDescription: " Remove one or more command using this syntax\n\n     jim --rm command_1 <command_2> ...\n\n Will remove the provided commands.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},
}
