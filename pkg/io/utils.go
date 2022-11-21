package io

import (
	"errors"
	"fmt"
	"io/ioutil"
	"jim/internal/constants"
	"jim/pkg/models"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func ExecutableFolder() string {
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(path)

	return filepath.Dir(path)
}

func Executable() string {
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(path)

	return path
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

func ReplaceCurrentFolderFlag(input string) string {
	return strings.Replace(input, constants.CURRENT_FOLDER_FLAG, CurrentFolder(), -1)
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

	var (
		editor string
		found  bool
	)

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

// get the command properties from the user inputs
// - if set it gets it from the args
// - otherwise it opens a temp file
func GetCommandValueFromArgs(args []string, command *models.Command) error {

	var (
		new_command_value string
		err               error
	)

	if len(args) == 1 {

		new_command_value, err = fileInput(command.Value)

		if err != nil {
			return err
		}

		new_command_value = strings.TrimSpace(new_command_value)

		if new_command_value == "" {
			return errors.New("the command cannot be empty")
		}
	} else {
		new_command_value = args[1]
	}

	command.Value = ReplaceCurrentFolderFlag(new_command_value)

	return nil
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

func DetachedCmd(arg ...string) (*exec.Cmd, error) {

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
		"'" + arg[0] + "' " + strings.Join(arg[1:], " ") + "& disown",
	)

	return c, err

}

func InterceptStdout(callback func()) string {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	callback()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return strings.TrimSpace(string(out))
}

func RunMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func IsRunningAsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

	return err == nil
}
