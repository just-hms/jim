package actions

import (
	"fmt"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestRemove(t *testing.T) {

	rainbow.Blank()

	mockResponse := ""

	utils.InterceptStdout(func() {

		// clear all commands
		Clear.Value([]string{"--force"})

		args := []string{"print", "echo 1"}

		// add a command
		Add.Value(args)

		// remove it
		if !Remove.ArgumentsCheck(args[:1]) {
			fmt.Println("wrong format")
			return
		}

		Remove.Value(args[:1])
	})

	// check that it has been removed
	responseData := utils.InterceptStdout(func() {
		List.Value([]string{})
	})

	assert.Equal(t, responseData, mockResponse)
}
