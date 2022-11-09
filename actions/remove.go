package actions

import (
	"jim/models"
	"jim/utils"

	"github.com/tidwall/buntdb"
)

var Remove = &Action{
	Value: func(args []string) {

		for _, arg := range args {

			command := models.Command{}

			if err := FindCommandByName(arg, &command); err != nil {
				return
			}

			err := models.DB().Update(func(tx *buntdb.Tx) error {

				// delete all sessions
				var delkeys []string

				tx.AscendKeys("command:session:"+command.Name+":*", func(k, v string) bool {
					delkeys = append(delkeys, k)
					return true // continue
				})

				for _, k := range delkeys {
					if _, err := tx.Delete(k); err != nil {
						return err
					}
				}

				_, err := tx.Delete("command:" + command.Name)
				return err
			})

			if err != nil {
				utils.Alertf("error while deleting\n")
			}
		}
	},
	Description:     "remove one or more specified command",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},
}
