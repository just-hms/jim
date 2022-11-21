package actions

import (
	"fmt"
	"jim/pkg/io"
	"jim/pkg/rainbow"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestHelp(t *testing.T) {

	rainbow.Blank()

	mockResponses := []string{
		"wrong format",
		strings.TrimSpace(Add.HelpDescription),
	}

	argss := [][]string{
		{"", "", ""},
		{"--add"},
	}

	// create three commands

	for i := 0; i < len(argss); i++ {

		responseData := io.InterceptStdout(func() {

			if !Help.ArgumentsCheck(argss[i]) {
				fmt.Println("wrong format")
				return
			}

			Help.Value(argss[i])
		})

		assert.Equal(t, responseData, mockResponses[i])
	}

}
