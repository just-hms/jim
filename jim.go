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

	// if no action  was found call the run command
	if action == nil {

		actions.Run.Value([]string{
			command,
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
		utils.Alertf("wrong format!!!")
		return
	}

	// check if the command needs to be launched in background

	bg := os.Getenv("JIM_ENV")

	// if yes and it is not launch it in backgorund
	if action.BackGround && (bg == "" || bg != "background") {

		executable, _ := os.Executable()

		c, err := utils.DetachedCrossCmd(
			executable,
			command,
			strings.Join(args[0:], " "),
		)

		if err != nil {
			utils.Alertf(err.Error())
			return
		}

		c.Env = append(os.Environ(), "JIM_ENV=background")
		c.Run()
		return
	}

	// call the action
	action.Value(args)

}
