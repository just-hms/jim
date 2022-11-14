package test

import (
	"jim/actions"
	"testing"

	"github.com/go-playground/assert"
)

func TestWatchShow(t *testing.T) {

	mockResponse := ""

	// clearr and add a command and watch it

	interceptStdout(func() {
		actions.Clear.Value([]string{"--force"})
		actions.Add.Value([]string{"print", "echo 1"})
		actions.Watch.Value([]string{"print"})
	})

	responseData := interceptStdout(func() {
		actions.Show.Value([]string{})
	})

	// check if the wacth result is in the db
	assert.NotEqual(t, responseData, mockResponse)

}
