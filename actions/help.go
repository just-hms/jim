package actions

import (
	"fmt"
	"jim/utils"
	"sort"
)

var Help = &Action{

	Value: func(args []string) {

		fmt.Println("The jim command line utility enables running long commands with one word")
		fmt.Println("")

		fmt.Printf("usage:\n")

		utils.Titlef("           jim [%s<action>] [<arguments>]\n", utils.ACTION_PREFIX)
		utils.Commentf("           to edit your commands\n")
		fmt.Println("")

		utils.Titlef("           jim %srun command\n", utils.ACTION_PREFIX)
		fmt.Println("      or")

		utils.Titlef("           jim command\n")
		utils.Commentf("           to launch a command\n")
		fmt.Println("")

		fmt.Println("The following Actions are available")

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
		utils.Titlef("\n           jim %sadd kek \"code %s\"\n", utils.ACTION_PREFIX, utils.CURRENT_FOLDER_FLAG)
		utils.Titlef("           jim %sls kek\n\n", utils.ACTION_PREFIX)
		fmt.Printf("and see what happens\n\n")

		fmt.Println("For more details on a specific action, pass it the help argument.")

	},
	Description:  "list of all 	 Actions and their description",
	ArgumentsLen: 0,
}