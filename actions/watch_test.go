package actions

import (
	"jim/models"
	"jim/rainbow"
	"jim/utils"
	"testing"

	"github.com/go-playground/assert"
)

func TestWatch(t *testing.T) {

	rainbow.Blank()

	mockResponse := ""

	// clear and add a command and watch it

	command := models.Command{
		Name:  "strange_name",
		Value: "echo lolz",
	}

	command.Save()

	Watch.Value([]string{command.Name})

	responseData := utils.InterceptStdout(func() {
		Show.Value([]string{command.Name})
	})

	// check if the wacth result is in the db
	assert.NotEqual(t, responseData, mockResponse)

}
