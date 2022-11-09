package actions

import (
	"jim/models"
	"jim/utils"
	"strings"

	"github.com/tidwall/buntdb"
)

var Rename = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		to_rename_to := args[1]

		getErr := models.DB().View(func(tx *buntdb.Tx) error {
			_, err := tx.Get("command:" + to_rename_to)
			return err
		})

		if getErr == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		setErr := models.DB().Update(func(tx *buntdb.Tx) error {

			var rnKeys []string

			// iterate over the sessions

			tx.AscendKeys("session:"+command.Name+":*", func(k, v string) bool {
				rnKeys = append(rnKeys, k)
				return true // continue
			})

			for _, k := range rnKeys {

				// delete the session
				val, delErr := tx.Delete(k)

				if delErr != nil {
					return delErr
				}

				// set the session value to the new key

				_, _, setErr := tx.Set(
					strings.ReplaceAll(k, command.Name, to_rename_to),
					val, // the old session value
					nil,
				)

				if setErr != nil {
					return setErr
				}

			}

			// delete the old commands key and set the new one
			val, err := tx.Delete("command:" + command.Name)

			if err != nil {
				return err
			}

			_, _, setErr := tx.Set("command:"+to_rename_to, val, nil)
			return setErr
		})

		if setErr != nil {
			utils.Alertf("error renaming the command\n")
			return
		}

	},
	Description:     "rename a command",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 2
	},
}
