package actions

import (
	"fmt"
	"jim/utils"
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

		utils.InterceptStdout(func() {
			if !Add.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			Add.Value(argss[i])
		})
	}

	responseData := utils.InterceptStdout(func() {
		List.Value([]string{})
	})

	assert.NotEqual(t, responseData, mockResponse)

}
