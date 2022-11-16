package actions

import (
	"errors"
	"jim/internal/constants"
	"jim/pkg/levenshtein"
	"jim/pkg/models"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"os"
	"strings"
)

type Action struct {
	Value               func([]string)
	Description         string
	ArgumentsCheck      func([]string) bool
	HelpDescription     string
	BackgroundSubAction func([]string)
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
		newActions[constants.ACTION_PREFIX+k] = Actions[k]
	}

	Actions = newActions

}

// uses the https://en.wikipedia.org/wiki/Levenshtein_distance
// returns error if no command was found

func FindCommandByName(name string, command *models.Command) error {

	if err := models.GetCommandByName(command, name); err == nil {
		return nil
	}

	// if no result was found try with similiar names

	var commands []models.Command

	if err := models.GetCommands("", &commands); err != nil {
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

	if max_lev_rateo >= constants.MIN_ACCEPTABLE_LEV_RATEO {
		return nil
	}

	rainbow.Warningf("did you mean %s? Type y or N\n", command.Name)

	if utils.ReadChar() == 'y' {
		return nil
	}

	return errors.New("command not found")

}

func ContinueInBackground(command models.Command, params string) {

	executable, _ := os.Executable()

	action := constants.BG_ACTION_PREFIX + strings.Replace(os.Args[1], constants.ACTION_PREFIX, "", -1)

	c, _ := utils.DetachedCmd(
		executable,
		action,
		command.Name,
		params,
	)

	c.Stderr = os.Stderr
	c.Run()
}

func TakeUp(args []string) (models.Command, string, error) {

	var command models.Command

	if err := models.GetCommandByName(&command, args[0]); err != nil {
		return command, "", err
	}

	return command, strings.Join(args[1:], " "), nil
}
