package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func ExecutableFolder() string {
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(path)

	return filepath.Dir(path)
}

func CurrentFolder() string {
	path, _ := os.Getwd()
	return path
}

func ReadChar() rune {

	b := make([]byte, 1)
	os.Stdin.Read(b)

	if rune(string(b)[0]) == 13 {
		return 'y'
	}

	return rune(string(b)[0])
}

var Alertf = color.New(color.FgRed).PrintfFunc()
var Titlef = color.New(color.FgHiWhite, color.Bold).PrintfFunc()
var Commentf = color.New(color.FgHiBlack, color.Bold).PrintfFunc()
var Warningf = color.New().PrintfFunc()

func ReplaceCurrentFolderFlag(input string) string {
	return strings.Replace(input, CURRENT_FOLDER_FLAG, CurrentFolder(), -1)
}

func fileInput(file_default_content string) (string, error) {

	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "command")

	if tmpFileErr != nil {
		return "", errors.New("error while creating tempFile")
	}

	// set the file content to file_default_content
	tmpFile.WriteString(file_default_content)
	tmpFile.Close()

	// get the default editor
	var editor string
	var found bool

	if runtime.GOOS == "windows" {
		editor = "notepad"

	} else {

		editor, found = os.LookupEnv("VISUAL")

		// this doesn't work
		if !found {
			editor, found = os.LookupEnv("EDITOR")
		}
		if !found {
			editor = "vi"
		}
	}

	// get the default editor path
	path, err := exec.LookPath(editor)

	if err != nil {
		return "", errors.New("text editor not found")
	}

	// edit the tmp file with the default editor
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()

	if err != nil {
		return "", errors.New("text editor start failed")
	}

	cmd.Wait()

	data, err := os.ReadFile(tmpFile.Name())

	if err != nil {
		return "", errors.New("error reading the text file output")
	}

	return string(data), nil

}

func GetCommandFromArgs(args []string, old_command string) (string, error) {

	var command_value string
	var err error

	if len(args) == 1 {

		command_value, err = fileInput(old_command)

		if err != nil {
			return "", err
		}

		command_value = strings.TrimSpace(command_value)

		if command_value == "" {
			return "", errors.New("the command cannot be empty")
		}
	} else {
		command_value = args[1]
	}

	command_value = ReplaceCurrentFolderFlag(command_value)

	return command_value, nil
}

func CrossCmd(arg ...string) (*exec.Cmd, error) {

	var c *exec.Cmd

	if runtime.GOOS == "windows" {

		c = exec.Command(
			"powershell",
			append([]string{"-c"}, arg...)...,
		)
	} else {

		shell, err := os.LookupEnv("SHELL")

		if !err {
			return c, errors.New("no shell found")
		}

		c = exec.Command(
			shell,
			append([]string{"-c"}, arg...)...,
		)
	}

	return c, nil

}

func DetachedCrossCmd(arg ...string) (*exec.Cmd, error) {

	var (
		c   *exec.Cmd
		err error
	)

	if runtime.GOOS == "windows" {
		c, err = CrossCmd(
			"Invoke-Expression",

			"'cmd /c start powershell -windowstyle hidden -c "+
				strings.Join(arg, " ")+
				"'",
		)
		return c, err
	}

	c, err = CrossCmd(
		strings.Join(arg, " "),
		"& disown",
	)

	return c, err

}
