package actions

import (
	"fmt"
	"jim/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestMod(t *testing.T) {

	mockResponse := "wrong format"
	args := []string{}

	// check if the args check works

	responseData := utils.InterceptStdout(func() {

		if !Mod.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}

		Mod.Value(args)

	})

	assert.Equal(t, responseData, mockResponse)
}
