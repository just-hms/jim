package io

import (
	"errors"
	"io/ioutil"
	"jim/internal/constants"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func FileInput(file_default_content string) (string, error) {

	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "command")

	if tmpFileErr != nil {
		return "", errors.New("error while creating tempFile")
	}

	// set the file content to file_default_content
	tmpFile.WriteString(file_default_content)
	tmpFile.Close()

	// get the default editor
	editor := GetDefaultTextEditor()

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
