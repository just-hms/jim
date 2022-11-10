package actions

import (
	"fmt"
	"jim/utils"
)

var Version = &Action{
	Value: func(args []string) {
		fmt.Println(utils.Version)
	},
	Description: "print the version of the executable",
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
	HelpDescription: " Show the version of the executable using this syntax\n\n     jim --version\n\n Will output the installed jim version ex: 'v1.0.1'.",
}
