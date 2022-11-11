package test

import (
	"fmt"
	"jim/actions"
	"jim/utils"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestVersion(t *testing.T) {

	var (
		mockResponse string
		responseData string
	)

	// correct argument test
	mockResponse = utils.Version

	responseData = getStdout(func() {
		args := []string{}

		if !actions.Version.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}
		actions.Version.Value(args)
	})

	assert.Equal(t, strings.TrimSpace(responseData), mockResponse)

	// wrong argument test

	mockResponse = "wrong format"

	responseData = getStdout(func() {
		args := []string{"wrong input"}

		if !actions.Version.ArgumentsCheck(args) {
			fmt.Println("wrong format")
			return
		}
		actions.Version.Value(args)
	})

	assert.Equal(t, strings.TrimSpace(responseData), mockResponse)

}
