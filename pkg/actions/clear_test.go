package actions

import (
	"jim/pkg/io"
	"jim/pkg/rainbow"
	"testing"

	"github.com/go-playground/assert"
)

func TestClear(t *testing.T) {

	rainbow.Blank()
	mockResponse := ""

	argss := [][]string{
		{"print1", "echo 1"},
		{"print2", "echo 2"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {
		io.InterceptStdout(func() {
			Add.Value(argss[i])
		})
	}

	// clear them and check that the output is ""

	responseData := io.InterceptStdout(func() {
		Clear.Value([]string{"--force"})
		List.Value([]string{})
	})

	assert.Equal(t, responseData, mockResponse)

}
