package actions

import (
	"errors"
	"jim/levenshtein"
	"jim/models"
	"jim/utils"

	"github.com/tidwall/buntdb"
)

type Action struct {
	Value           func([]string)
	Description     string
	ArgumentsCheck  func([]string) bool
	HelpDescription string
	BackgroundShit  func([]string)
}

var Actions = map[string]*Action{}

func init() {

	Actions = map[string]*Action{
		"ls":      List,
		"add":     Add,
		"mod":     Mod,
		"rm":      Remove,
		"clear":   Clear,
		"run":     Run,
		"rn":      Rename,
		"help":    Help,
		"watch":   Watch,
		"show":    Show,
		"version": Version,
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

	getErr := models.DB().View(func(tx *buntdb.Tx) error {
		var err error
		command.Value, err = tx.Get("command:" + name)
		command.Name = name // set the name to the key if found
		return err
	})

	if getErr == nil {
		return nil
	}

	// if no result was found try with similiar names

	var commands []models.Command

	if err := models.ListCommands("", &commands); err != nil {
		return errors.New("error retrieving the command")
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
