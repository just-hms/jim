package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"strings"
)

var List = &Action{
	Value: func(args []string) {

		var commands []models.Command

		filter := ""

		if len(args) >= 1 {
			filter = args[0]
		}

		if err := models.ListCommands(filter, &commands); err != nil {
			utils.Alertf("error retrieving the command\n")
			return
		}

		if len(commands) == 0 {
			return
		}

		utils.Titlef(" %-30s%-1s\n", "Command", "Value")

		multiline_tab := "\n                               "

		for _, v := range commands {

			trimmed := strings.TrimSpace(v.Value)

			// tab the lines after the first
			tabbed := strings.ReplaceAll(trimmed, "\n", multiline_tab)

			fmt.Printf(" %-30s%-1s\n", v.Name, tabbed)
		}

	},
	Description:     "list of all commands",
	HelpDescription: "wp",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0 || len(args) == 1
	},
}
