package actions

import (
	"fmt"
	"jim/pkg/io"
	"jim/pkg/rainbow"
	"testing"

	"github.com/go-playground/assert"
)

func TestList(t *testing.T) {

	rainbow.Blank()

	mockResponse := ""

	argss := [][]string{
		{"print1", "echo 1"},
		{"print2", "echo 2"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {

		io.InterceptStdout(func() {
			if !Add.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			Add.Value(argss[i])
		})
	}

	responseData := io.InterceptStdout(func() {
		List.Value([]string{})
	})

	assert.NotEqual(t, responseData, mockResponse)

}
