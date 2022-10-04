package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
	"strings"
)

var List = &Action{
	Value: func(args []string) {

		if len(args) != 0 && len(args) != 1 {
			utils.Alertf("wrong format!!!")
			return
		}

		commands := []models.Command{}
		models.DB().Table("commands").Scan(&commands)

		if len(commands) == 0 {
			fmt.Println("...")
			return
		}

		filtered := []models.Command{}

		if len(args) == 1 {
			for i := range commands {
				if strings.Contains(commands[i].Name, args[0]) {
					filtered = append(filtered, commands[i])
				}
			}
		} else {
			filtered = commands
		}

		utils.Titlef(" %-10s%-30s%-1s\n", "ID", "Name", "Comand")

		for _, v := range filtered {
			fmt.Printf(" %-10d%-30s%-1s\n", v.ID, v.Name, v.Value)
		}

	},
	Description:     "list of all commands",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
	HelpDescription: "wp",
}
