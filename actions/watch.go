package actions

import (
	"jim/models"
	"jim/utils"
	"strconv"

	"strings"
	"time"

	"github.com/tidwall/buntdb"
)

var Watch = &Action{
	Value: func(args []string) {

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			utils.Alertf("%s\n", err.Error())
			return
		}

		params := ""

		if len(args) == 2 {
			params = strings.Join(args[1:], " ")
		}

		utils.ContinueInBackGround(command, params)

	},
	Description:     "run a command in background and time it (user input and output don't work)",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) >= 1
	},

	BackgroundShit: func(args []string) {

		command, params, err := utils.TakeUp(args)

		if err != nil {
			return
		}

		c, err := utils.CrossCmd(command.Value, params)

		if err != nil {
			return
		}

		session := models.Session{
			Start:   time.Now(),
			Command: command.Name,
		}

		if err := c.Run(); err != nil {
			return
		}

		session.Elapsed = time.Since(session.Start)

		setErr := models.DB().Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(
				"session:command:"+command.Name+":"+strconv.FormatInt(session.Start.Unix(), 10),
				strconv.FormatInt(session.Elapsed.Milliseconds(), 10),
				nil,
			)
			return err
		})

		if setErr != nil {
			utils.Alertf("error adding the session\n")
			return
		}
	},
}
