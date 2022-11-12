package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestRename(t *testing.T) {

	mockResponse := ""

	interceptStdout(func() {

		// clear all commands
		actions.Clear.Value([]string{"--force"})

		add_args := []string{"print", "echo 1"}
		rn_args := []string{"print", "kek"}

		// add a command
		actions.Add.Value(add_args)

		// rename it
		if !actions.Rename.ArgumentsCheck(rn_args) {
			fmt.Println("wrong format")
			return
		}

		actions.Rename.Value(rn_args)
	})

	// check that it has been renamed

	responseData := interceptStdout(func() {
		actions.List.Value([]string{"kek"})
	})

	assert.NotEqual(t, responseData, mockResponse)
}
