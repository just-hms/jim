package actions

import (
	"jim/models"
	"jim/utils"

	"github.com/tidwall/buntdb"
)

var Mod = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		if err := utils.GetCommandFromArgs(args, &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		setErr := models.DB().Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set("command:"+command.Name, command.Value, nil)
			return err
		})

		if setErr != nil {
			utils.Alertf("error adding the command\n")
			return
		}
	},
	Description:     "modify a specified command",
	HelpDescription: " Modify a command using this syntax\n\n     jim --mod command\n\n Will open the command in your default editor and will let\n you modify it.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}
