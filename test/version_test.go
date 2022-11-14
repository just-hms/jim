package test

import (
	"fmt"
	"jim/actions"
	"jim/constants"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestVersion(t *testing.T) {

	var responseData string

	// correct argument test
	corrMockResponse := constants.Version
	errMockResponse := "wrong format"

	responseData = interceptStdout(func() {

		// correct input

		args := []string{}

		if !actions.Version.ArgumentsCheck(args) {
			fmt.Println(errMockResponse)
			return
		}
		actions.Version.Value(args)
	})

	assert.Equal(t, strings.TrimSpace(responseData), corrMockResponse)

	// wrong argument test

	responseData = interceptStdout(func() {

		// wrong input

		args := []string{""}

		if !actions.Version.ArgumentsCheck(args) {
			fmt.Println(errMockResponse)
			return
		}
		actions.Version.Value(args)
	})

	assert.Equal(t, responseData, errMockResponse)

}
