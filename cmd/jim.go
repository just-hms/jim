package main

import (
	"jim/actions"
	"jim/utils"

	"os"
	"strings"
)

func main() {

	// display help if no command was provided
	if len(os.Args) <= 1 {
		actions.Help.Value([]string{})
		return
	}

	// the input_command is the first argument
	input_command := os.Args[1]
	var args []string

	// the args are the other

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// get the action from the command
	action := actions.Actions[input_command]

	// if no action was found check if it's the background one
	if action == nil {

		action = actions.Actions[utils.ACTION_PREFIX+
			strings.Replace(
				input_command,
				utils.BG_PREFIX, "", -1,
			)]

		if action != nil {
			action.BackgroundShit(args)
			return
		}
	}

	// if no action  was found call the run command
	if action == nil {

		actions.Run.Value([]string{
			input_command,
			strings.Join(args, " "),
		})

		return
	}

	// if the user wrote: "jim --action --help"

	if len(args) == 1 && args[0] == utils.ACTION_PREFIX+"help" {
		println(action.HelpDescription)
		return
	}

	// check if the argument len is correct

	if !action.ArgumentsCheck(args) {
		utils.Alertf("wrong format!!!\n")
		return
	}

	// call the action
	action.Value(args)

}