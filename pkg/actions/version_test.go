package actions

import (
	"fmt"
	"jim/internal/constants"
	"jim/pkg/io"
	"jim/pkg/rainbow"
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

	responseData = io.InterceptStdout(func() {

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

	responseData = io.InterceptStdout(func() {

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
