package actions

import (
	"fmt"
	"jim/utils"
	"sort"

	"github.com/fatih/color"
)

var Help = &Action{
	Value: func(args []string) {

		fmt.Println("The jim command line utility enables running long commands with one word")
		fmt.Println("")

		fmt.Printf("usage: jim [%s<action>] [<arguments>]\n", utils.ACTION_PREFIX)
		fmt.Println("")

		fmt.Println("The following Actions are available")

		keys := make([]string, 0, len(Actions))
		for k := range Actions {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {

			fmt.Printf(" %s", color.HiWhiteString(key))

			for i := 0; i < 10-len(key); i++ {
				fmt.Print(" ")
			}

			fmt.Println(Actions[key].Description)
		}

		fmt.Println("")

		fmt.Println("For more details on a specific action, pass it the help argument.")

	},
	Description:  "list of all possible Actions and their descriprion",
	ArgumentsLen: 0,
}
