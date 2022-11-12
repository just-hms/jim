package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestRun(t *testing.T) {

	mockResponse := "wrong format"
	args := []string{}

	// check if the args check works

	responseData := interceptStdout(func() {

		if !actions.Run.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}

		actions.Run.Value(args)

	})

	assert.Equal(t, responseData, mockResponse)
}
