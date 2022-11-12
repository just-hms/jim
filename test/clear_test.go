package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestClear(t *testing.T) {

	mockResponse := ""

	argss := [][]string{
		{"print1", "echo 1"},
		{"print2", "echo 2"},
		{"print3", "echo 3"},
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

	// clear them and check that the output is ""

	responseData := getStdout(func() {
		actions.Clear.Value([]string{"--force"})
		actions.List.Value([]string{})
	})

	assert.Equal(t, responseData, mockResponse)

}
