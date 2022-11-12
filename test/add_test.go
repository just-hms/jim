package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestAdd(t *testing.T) {

	var responseData string

	mockResponses := []string{
		"wrong format",
		"",
		"a command named print already exists!!!",
	}

	argss := [][]string{
		{},
		{"print", "echo hello i'm jim"},
		{"print", ""},
	}

	for i := 0; i < len(argss); i++ {

		responseData = getStdout(func() {

			if !actions.Add.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			actions.Add.Value(argss[i])

		})

		assert.Equal(t, responseData, mockResponses[i])
	}

}
