package test

import (
	"fmt"
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestMod(t *testing.T) {

	mockResponse := "wrong format"
	args := []string{}

	// check if the args check works

	responseData := interceptStdout(func() {

		if !actions.Mod.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}

		actions.Mod.Value(args)

	})

	assert.Equal(t, responseData, mockResponse)
}
