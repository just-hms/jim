package actions

import (
	"fmt"
	"jim/models"
	"jim/utils"
)

var List = &Action{
	Value: func(args []string) {

		commands := []models.Command{}
		models.DB().Table("commands").Scan(&commands)

		if len(commands) == 0 {
			fmt.Println("...")
			return
		}

		utils.Titlef(" %-10s%-30s%-1s\n", "ID", "Name", "Comand")

		for _, v := range commands {
			fmt.Printf(" %-10d%-30s%-1s\n", v.ID, v.Name, v.Value)
		}

	},
	Description:     "list of all commands",
	ArgumentsLen:    utils.CUSTOM_ARGUMENTS_LEN,
	HelpDescription: "wp",
}
