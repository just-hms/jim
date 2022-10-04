package actions

import (
	"errors"
	"fmt"
	"jim/levenshtein"
	"jim/models"
	"jim/utils"
)

type Action struct {
	Value           func([]string)
	Description     string
	ArgumentsLen    int
	HelpDescription string
}

var Actions = map[string]*Action{}

func init() {

	Actions = map[string]*Action{
		"ls":   List,
		"add":  Add,
		"mod":  Mod,
		"modi": ModById,
		"rm":   Remove,
		"rmi":  RemoveById,
		"cls":  Clear,
		"run":  Run,
		"rn":   Rename,
		"help": Help,
	}

	var newActions = map[string]*Action{}

	for k := range Actions {
		newActions[utils.ACTION_PREFIX+k] = Actions[k]
	}

	Actions = newActions

}

func FindCommandByName(name string, command *models.Command) error {

	if err := models.DB().Where("name = ?", name).First(command).Error; err == nil {
		return nil
	}

	// if no result was found try with similiar names

	commands := []models.Command{}
	models.DB().Find(&commands)

	if len(commands) == 0 {
		fmt.Println("...")
		return errors.New("commands empty")
	}

	max_lev_rateo := 0.0

	for i, item := range commands {

		if len(item.Name) == 0 {
			continue
		}

		item_lev :=
			float64((len(item.Name) - levenshtein.Levenshtein(item.Name, name))) /
				float64(len(item.Name))

		if item_lev > max_lev_rateo {
			max_lev_rateo = item_lev
			*command = commands[i]
		}
	}

	if max_lev_rateo == 0 {
		fmt.Println("...")
		return errors.New("nothing similiar found")

	}

	if max_lev_rateo >= 0.80 {
		return nil
	}

	utils.Warningf("did you mean %s? Type y or N\n", command.Name)

	if utils.ReadChar() == 'y' {
		return nil
	}

	return errors.New("not found")

}

func ArgumentsLenCorresponds(action *Action, args []string) bool {

	return action.ArgumentsLen == len(args) || action.ArgumentsLen == utils.CUSTOM_ARGUMENTS_LEN
}
