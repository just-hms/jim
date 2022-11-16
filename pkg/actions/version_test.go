package actions

import (
	"fmt"
	"jim/internal/constants"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestVersion(t *testing.T) {

	rainbow.Blank()

	var responseData string

	// correct argument test
	corrMockResponse := constants.Version
	errMockResponse := "wrong format"

	responseData = utils.InterceptStdout(func() {

		// correct input

		args := []string{}

		if !Version.ArgumentsCheck(args) {
			fmt.Println(errMockResponse)
			return
		}
		Version.Value(args)
	})

	assert.Equal(t, strings.TrimSpace(responseData), corrMockResponse)

	// wrong argument test

	responseData = utils.InterceptStdout(func() {

		// wrong input

		args := []string{""}

		if !Version.ArgumentsCheck(args) {
			fmt.Println(errMockResponse)
			return
		}
		Version.Value(args)
	})

	assert.Equal(t, responseData, errMockResponse)

}
