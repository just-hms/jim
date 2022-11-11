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

		var sessions []models.Session

		filter := ""

		if len(args) >= 1 {
			filter = args[0]
		}

		if err := models.GetSessions(filter, &sessions); err != nil {
			utils.Alertf("error retrieving the sessions\n")
			return
		}

		if len(sessions) == 0 {
			return
		}

		utils.Titlef(" %-10s%-30s%-1s\n", "Command", "Date", "Duration")

		total := 0
		for _, s := range sessions {
			startDate := s.Start.Format("2006-01-02 15:04:05")
			fmt.Printf(" %-10s%-30s%-1s\n", s.Command, startDate, s.Elapsed)
			total += int(s.Elapsed)
		}

		fmt.Printf("\n Total := %s\n", time.Duration(total))
	},
	Description: "show the watching results",
	ArgumentsCheck: func(args []string) bool {
		return true
	},
	HelpDescription: " Show a list of all the --watch results using this syntax\n\n     jim --show <filter>\n\n Will list all of the commands' sessions. Filtering them with\n the provided filter.\n A session is created when a command is launched with --watch.",
	BackgroundShit: func([]string) {
	},
}
