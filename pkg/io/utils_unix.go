//go:build darwin || linux

package io

import (
	"errors"
	"jim/pkg/rainbow"
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

func RequireAdmin() {
	if os.Getuid() != 0 {
		rainbow.Alertf("permission denied\n")
		os.Exit(0)
	}
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

func ConfigFolder() string {
	configFolder, _ := os.LookupEnv("HOME")
	configFolder = filepath.Join(configFolder, "/.local/share/jim")

	return configFolder
}
