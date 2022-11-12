package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestList(t *testing.T) {

	mockResponse := ""

	argss := [][]string{
		{"print1", "echo 1"},
		{"print2", "echo 2"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {

		getStdout(func() {
			if !actions.Add.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			actions.Add.Value(argss[i])
		})
	}

	responseData := getStdout(func() {
		actions.List.Value([]string{})
	})

	assert.NotEqual(t, responseData, mockResponse)

}
