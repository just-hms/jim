package actions

import (
	"jim/models"
	"jim/utils"

	"github.com/tidwall/buntdb"
)

var Add = &Action{
	Value: func(args []string) {

		command := models.Command{
			Name: args[0],
		}

		getErr := models.DB().View(func(tx *buntdb.Tx) error {
			_, err := tx.Get("command:" + command.Name)
			return err
		})

		if getErr == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
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
	Description:     "add a command",
	HelpDescription: " add a command using this syntax\n\n     jim --add command <value>\n\n If no value is specified jim will open your default editor and\n will let you insert a set of instruction in a temporary file.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}
