//go:build darwin || linux

package io

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DetachedCmd(arg ...string) (*exec.Cmd, error) {
	return CrossCmd(
		"'" + arg[0] + "' " + strings.Join(arg[1:], " ") + "& disown",
	)
}

func CrossCmd(arg ...string) (*exec.Cmd, error) {

	shell, err := os.LookupEnv("SHELL")

	if !err {
		return nil, errors.New("no shell found")
	}

	c := exec.Command(
		shell,
		append([]string{"-c"}, arg...)...,
	)

	return c, nil

}

func GetDefaultTextEditor() string {

	editor, found := os.LookupEnv("VISUAL")

	if !found {
		editor, found = os.LookupEnv("EDITOR")
	}

	if !found {
		editor = "vi"
	}

	return editor
}

func RequireAdmin() error {
	if os.Getuid() != 0 {
		return errors.New("permission denied")
		os.Exit(0)
	}

	return nil
}

func ConfigFolder() string {

	configFolder := os.Getenv("HOME")

	// get the config folder of the sudo user
	if isRunningAsAdmin() {

		sudo_user := os.Getenv("SUDO_USER")

		if sudo_user != "" {
			configFolder, _ = getUserHome(sudo_user)
		}
	}

	configFolder = filepath.Join(configFolder, ".local/share/jim")
	return configFolder

}

func isRunningAsAdmin() bool {
	return os.Getuid() == 0
}

func getUserHome(username string) (string, error) {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		return "", errors.New("error accessing config folder")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, username) {
			return strings.Split(line, ":")[5], nil
		}
	}

	return "", errors.New("user not found")

}
