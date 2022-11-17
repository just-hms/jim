package actions

import (
	"fmt"
	"jim/pkg/models"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestWatch(t *testing.T) {

	rainbow.Blank()

	command := models.Command{
		Name:  "to_watch",
		Value: "echo 1",
	}

	command.Save()

	// correct test
	correctMockResponse := "jim is launching > [powershell -c echo 1 ]\r"
	correctArgs := []string{"to_watch"}

	correctResponseData := utils.InterceptStdout(func() {

		if !Watch.ArgumentsCheck(correctArgs) {
			fmt.Println("wrong format")
			return
		}

		Watch.Value(correctArgs)
		// Show.Value([]string{})
	})

	assert.Equal(t, strings.TrimSpace(correctResponseData), strings.TrimSpace(correctMockResponse))

	// wrong test
	wrongMockResponse := "wrong format"
	wrongArgs := []string{}

	// check if the args check works

	wrongResponseData := utils.InterceptStdout(func() {

		if !Watch.ArgumentsCheck(wrongArgs) {
			fmt.Println("wrong format")
			return
		}

		Watch.Value(wrongArgs)

	})

	assert.Equal(t, wrongResponseData, wrongMockResponse)
}
