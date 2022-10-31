package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"time"
)

// format:
// jim --show <command> --from date --to date

// todo:
// get inputs
// - show id if no command is specified
// - always show from and to
// - add toady as a possible value

var Show = &Action{
	Value: func(args []string) {

		sessions := []models.Session{}

		if len(args) >= 1 && args[0] != utils.ACTION_PREFIX+"from" && args[0] != utils.ACTION_PREFIX+"to" {

			command := models.Command{}
			if err := FindCommandByName(args[0], &command); err != nil {
				return
			}

			models.Eager().Where("command_id = ?", command.ID).Find(&sessions)
		} else {
			models.Eager().Find(&sessions)
		}

		utils.Titlef(" %-10s%-30s%-1s\n", "Command", "Date", "Duration")

		total := 0

		for _, s := range sessions {
			startDate := s.CreatedAt.Format("2006-01-02 15:04:05")

			fmt.Printf(" %-10s%-30s%-1s\n", s.Command.Name, startDate, s.Elapsed)

			total += int(s.Elapsed)
		}

		fmt.Printf("\n Total := %s\n", time.Duration(total))

	},
	Description:     "show the watching results",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return true
	},
}
