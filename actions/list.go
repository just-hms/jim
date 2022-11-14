package actions

import (
	"fmt"
	"jim/models"
	"jim/rainbow"
	"strings"
)

var List = &Action{
	Value: func(args []string) {

		var commands []models.Command

		filter := ""

		if len(args) >= 1 {
			filter = args[0]
		}

		if err := models.GetCommands(filter, &commands); err != nil {
			rainbow.Alertf("error retrieving the command\n")
			return
		}

		if len(commands) == 0 {
			return
		}

		rainbow.Titlef(" %-30s%-1s\n", "Command", "Value")

		multiline_tab := "\n                               "

		for _, v := range commands {

			trimmed := strings.TrimSpace(v.Value)

			// tab the lines after the first
			tabbed := strings.ReplaceAll(trimmed, "\n", multiline_tab)

			fmt.Printf(" %-30s%-1s\n", v.Name, tabbed)
		}

	},
	Description: "list of all commands",

	HelpDescription: " List of all the available commands using this syntax\n\n     jim --ls <filter>\n\n Will list all the available commands, filtering them with\n the provided filter.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0 || len(args) == 1
	},
}
