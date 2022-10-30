package main

import (
	"jim/actions"
	"jim/models"
	"jim/utils"

	"os"
	"strings"
)

func init() {
	models.Build()
}

func main() {

	// display help if no command was provided
	if len(os.Args) <= 1 {
		actions.Help.Value([]string{})
		return
	}

	// the command is the first argument
	command := os.Args[1]
	var args []string

	// the args are the other

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// get the action from the command
	action := actions.Actions[command]

	if action == nil {

		// if no action  was found call the run command
		actions.Run.Value([]string{
			command,
			strings.Join(args, " "),
		})
		return
	}

	// if the user wrote: "jim --add --help"
	if len(args) == 1 && args[0] == utils.ACTION_PREFIX+"help" {
		println(action.HelpDescription)
		return
	}

	// check if the argument len is correct

	if !actions.ArgumentsLenCorresponds(action, args) {
		utils.Alertf("wrong format!!!")
		return
	}

	action.Value(args)

}
