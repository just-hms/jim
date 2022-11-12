package test

import (
	"fmt"
	"jim/actions"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestHelp(t *testing.T) {

	mockResponses := []string{
		"wrong format",
		strings.TrimSpace(actions.Add.HelpDescription),
	}

	argss := [][]string{
		{"", "", ""},
		{"--add"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {

		responseData := interceptStdout(func() {

			if !actions.Help.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			actions.Help.Value(argss[i])
		})

		assert.Equal(t, responseData, mockResponses[i])
	}

}
