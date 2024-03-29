package actions

import (
	"fmt"
	"jim/pkg/io"
	"jim/pkg/rainbow"
	"testing"

	"github.com/go-playground/assert"
)

func TestAdd(t *testing.T) {

	rainbow.Blank()

	var responseData string

	mockResponses := []string{
		"wrong format",
		"",
		"a command named print already exists!!!",
	}

	argss := [][]string{
		{},
		{"print", "echo hello i'm jim"},
		{"print", ""},
	}

	for i := 0; i < len(argss); i++ {

		responseData = io.InterceptStdout(func() {

			if !Add.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			Add.Value(argss[i])

		})

		assert.Equal(t, responseData, mockResponses[i])
	}

}
