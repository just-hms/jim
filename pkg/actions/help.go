package actions

import (
	"fmt"
	"jim/internal/constants"
	"jim/pkg/rainbow"
	"sort"
)

var Help = &Action{

	Value: func(args []string) {

		// if an action is passed print it's help function
		if len(args) == 1 {
			action := Actions[args[0]]
			if action == nil || args[0] == constants.ACTION_PREFIX+"help" {
				rainbow.Alertf("action not found\n")
				return
			}

			fmt.Println(action.HelpDescription)
			return
		}

		// otherwise print the help text

		fmt.Print(
			"The jim command line utility enables running long commands with one word\n\n",
			"usage:\n",
		)

		rainbow.Titlef("           jim <%saction> <arguments>\n", constants.ACTION_PREFIX)
		rainbow.Commentf("           to manage your commands\n\n")

		rainbow.Titlef("           jim <%srun> command\n", constants.ACTION_PREFIX)
		rainbow.Commentf("           to launch a command\n\n")

		fmt.Println("The following actions are available")

		keys := make([]string, 0, len(Actions))
		for k := range Actions {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {

			rainbow.Titlef(" %s", key)

			for i := 0; i < 10-len(key); i++ {
				fmt.Print(" ")
			}

			fmt.Println(Actions[key].Description)
		}

		fmt.Println("")

		fmt.Printf("jim will change %s into the current path, try typing:\n", constants.CURRENT_FOLDER_FLAG)
		rainbow.Titlef("\n           jim %sadd . \"echo %s\"\n", constants.ACTION_PREFIX, constants.CURRENT_FOLDER_FLAG)
		rainbow.Titlef("           jim .\n\n")
		fmt.Printf("and see what happens\n\n")

		fmt.Println("For more details on a specific action, pass it the help argument.")

	},
	Description: "list of all actions and their description",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0 || len(args) == 1
	},
}
