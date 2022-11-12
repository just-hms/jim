package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestRemove(t *testing.T) {

	mockResponse := ""

	interceptStdout(func() {

		// clear all commands
		actions.Clear.Value([]string{"--force"})

		args := []string{"print", "echo 1"}

		// add a command
		actions.Add.Value(args)

		// remove it
		if !actions.Remove.ArgumentsCheck(args[:1]) {
			fmt.Println("wrong format")
			return
		}

		actions.Remove.Value(args[:1])
	})

	// check that it has been removed
	responseData := interceptStdout(func() {
		actions.List.Value([]string{})
	})

	assert.Equal(t, responseData, mockResponse)
}
