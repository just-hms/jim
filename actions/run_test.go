package actions

import (
	"fmt"
	"jim/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestRun(t *testing.T) {

	mockResponse := "wrong format"
	args := []string{}

	// check if the args check works

	responseData := utils.InterceptStdout(func() {

		if !Run.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}

		Run.Value(args)

	})

	assert.Equal(t, responseData, mockResponse)
}