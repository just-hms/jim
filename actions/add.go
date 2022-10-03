package actions

import (
	"jim/models"
	"jim/utils"
	"strings"
	"time"
)

var Add = &Action{
	Value: func(args []string) {

		to_search := models.Command{}

		if err := models.DB().Where("name = ?", args[0]).First(&to_search).Error; err == nil {
			utils.Alertf("a command named %s already exists!!!\n", args[0])
			return
		}

		args[1] = strings.Replace(args[1], utils.CURRENT_FOLDER_FLAG, utils.CurrentFolder(), -1)

		command := models.Command{
			Name:        args[0],
			Value:       args[1],
			LastTouched: time.Now(),
		}

		models.DB().Create(&command)

	},
	Description:     "add a command",
	HelpDescription: "wp",
	ArgumentsLen:    2,
}
