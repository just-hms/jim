package actions

import (
	"fmt"
	"io/ioutil"
	"jim/internal/constants"
	"jim/pkg/io"
	"jim/pkg/models"
	"jim/pkg/rainbow"
	"jim/pkg/test"
	"jim/pkg/tokenize"
	"os"
	"os/exec"
	"strings"
)

var Run = &Action{
	Value: func(args []string) {

		if len(args) != 1 && len(args) != 2 {
			rainbow.Alertf("wrong format!!!\n")
			return
		}

		command := models.Command{}

		if err := FindCommandByName(args[0], &command); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		if len(args) == 2 {
			RunCommand(command, strings.Join(args[1:], " "))
			return
		}

		RunCommand(command, "")
	},
	Description:     "run a command (not required)",
	HelpDescription: " Run a command using this syntax\n\n     jim <--run> command\n\n Will run the specified command in your default shell.\n --run can be omitted.",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1 || len(args) == 2
	},
}

func RunCommand(command models.Command, args string) {

	var (
		c   *exec.Cmd
		err error
	)

	if strings.TrimSpace(args) == "" {
		command.Value, err = tokenize.Tokenize(command.Value, []string{})
	} else {
		command.Value, err = tokenize.Tokenize(command.Value, strings.Split(args, ","))
	}

	if err != nil {
		rainbow.Alertf("%s\n", err.Error())
		return
	}

	if strings.HasPrefix(command.Value, constants.SHEBANG_PREFIX) {

		lines := strings.Split(command.Value, "\n")

		if len(lines) < 1 {
			return
		}

		tmpDir := os.TempDir()
		tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "command")

		if tmpFileErr != nil {
			return
		}

		exe := strings.TrimSpace(strings.Split(lines[0], constants.SHEBANG_PREFIX)[1])
		value := strings.Join(lines[1:], "\n")

		// set the file content to file_default_content
		tmpFile.WriteString(value)

		c = exec.Command(
			exe,
			tmpFile.Name(),
		)

		tmpFile.Close()

	} else {
		c, err = io.CrossCmd(command.Value)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}

	if test.IsTesting() {
		return
	}

	if err != nil {
		rainbow.Alertf("%s\n", err.Error())
		return
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		rainbow.Alertf("%s\n", err.Error())
	}
}
