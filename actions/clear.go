package actions

import (
	"jim/models"
	"jim/utils"

	"github.com/tidwall/buntdb"
)

var Clear = &Action{
	Value: func(args []string) {

		utils.Alertf("clear all commands is not reversible, are you sure? Type y or N\n")

		if utils.ReadChar() != 'y' {
			return
		}

		models.DB().Update(func(tx *buntdb.Tx) error {
			tx.DeleteAll()
			return nil
		})

	},
	Description:     "clear all commands",
	HelpDescription: " Clear all commands using this syntax\n\n     jim --clear\n\n Will remove all commands.",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
}
