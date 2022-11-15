package actions

import (
	"jim/rainbow"
	"jim/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestShow(t *testing.T) {

	rainbow.Blank()

	mockResponse := ""

	// clearr and add a command and watch it

	utils.InterceptStdout(func() {
		Clear.Value([]string{"--force"})
		Add.Value([]string{"print", "echo 1"})
		Watch.Value([]string{"print"})
	})

	responseData := utils.InterceptStdout(func() {
		Show.Value([]string{})
	})

	// check if the wacth result is in the db
	assert.NotEqual(t, responseData, mockResponse)

}
