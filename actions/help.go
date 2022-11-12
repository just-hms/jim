package actions

import (
	"fmt"
	"jim/utils"
	"sort"
)

var Help = &Action{

	Value: func(args []string) {

		// if an action is passed print it's help function
		if len(args) == 1 {
			action := Actions[args[0]]
			if action == nil || args[0] == utils.ACTION_PREFIX+"help" {
				utils.Alertf("action not found\n")
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

		utils.Titlef("           jim <%saction> <arguments>\n", utils.ACTION_PREFIX)
		utils.Commentf("           to manage your commands\n\n")

		utils.Titlef("           jim <%srun> command\n", utils.ACTION_PREFIX)
		utils.Commentf("           to launch a command\n\n")

		fmt.Println("The following actions are available")

		keys := make([]string, 0, len(Actions))
		for k := range Actions {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {

			utils.Titlef(" %s", key)

			for i := 0; i < 10-len(key); i++ {
				fmt.Print(" ")
			}

			fmt.Println(Actions[key].Description)
		}

		fmt.Println("")

		fmt.Printf("jim will change %s into the current path, try typing:\n", utils.CURRENT_FOLDER_FLAG)
		utils.Titlef("\n           jim %sadd . \"echo %s\"\n", utils.ACTION_PREFIX, utils.CURRENT_FOLDER_FLAG)
		utils.Titlef("           jim .\n\n")
		fmt.Printf("and see what happens\n\n")

		fmt.Println("For more details on a specific action, pass it the help argument.")

	},
	Description: "list of all actions and their description",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0 || len(args) == 1
	},
}
