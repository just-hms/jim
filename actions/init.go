package actions

import (
	"errors"
	"jim/levenshtein"
	"jim/models"
	"jim/utils"
)

type Action struct {
	Value           func([]string)
	Description     string
	ArgumentsCheck  func([]string) bool
	HelpDescription string
}

var Actions = map[string]*Action{}

func init() {

	Actions = map[string]*Action{
		"ls":    List,
		"add":   Add,
		"mod":   Mod,
		"modi":  ModById,
		"rm":    Remove,
		"rmi":   RemoveById,
		"clear": Clear,
		"run":   Run,
		"rn":    Rename,
		"help":  Help,
		"watch": WatchHelper,
		"show":  Show,

		"bg-watch": Watch,
	}

	// add prefix to actions hook

	var newActions = map[string]*Action{}

	for k := range Actions {
		newActions[utils.ACTION_PREFIX+k] = Actions[k]
	}

	Actions = newActions

}

// returns error if no command was found

func FindCommandByName(name string, command *models.Command) error {

	if err := models.DB().Where("name = ?", name).First(command).Error; err == nil {
		return nil
	}

	// if no result was found try with similiar names

	commands := []models.Command{}
	models.DB().Find(&commands)

	if len(commands) == 0 {
		return errors.New("there is no command yet")
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
		return errors.New("nothing similiar found")

	}

	if max_lev_rateo >= utils.MIN_ACCEPTABLE_LEV_RATEO {
		return nil
	}

	utils.Warningf("did you mean %s? Type y or N\n", command.Name)

	if utils.ReadChar() == 'y' {
		return nil
	}

	return errors.New("command not found")

}
