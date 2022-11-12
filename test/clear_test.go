package test

import (
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestClear(t *testing.T) {

	mockResponse := ""

	argss := [][]string{
		{"print1", "echo 1"},
		{"print2", "echo 2"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {
		interceptStdout(func() {
			actions.Add.Value(argss[i])
		})
	}

	// clear them and check that the output is ""

	responseData := interceptStdout(func() {
		actions.Clear.Value([]string{"--force"})
		actions.List.Value([]string{})
	})

	assert.Equal(t, responseData, mockResponse)

}
