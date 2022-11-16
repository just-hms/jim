package actions

import (
	"fmt"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestRename(t *testing.T) {

	rainbow.Blank()

	mockResponse := ""

	utils.InterceptStdout(func() {

		// clear all commands
		Clear.Value([]string{"--force"})

		add_args := []string{"print", "echo 1"}
		rn_args := []string{"print", "kek"}

		// add a command
		Add.Value(add_args)

		// rename it
		if !Rename.ArgumentsCheck(rn_args) {
			fmt.Println("wrong format")
			return
		}

		Rename.Value(rn_args)
	})

	// check that it has been renamed

	responseData := utils.InterceptStdout(func() {
		List.Value([]string{"kek"})
	})

	assert.NotEqual(t, responseData, mockResponse)
}
