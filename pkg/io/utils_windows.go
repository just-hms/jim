package io

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func runMeElevated() {
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

func isRunningAsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

	return err == nil
}

func DetachedCmd(arg ...string) (*exec.Cmd, error) {

	return CrossCmd(
		"Invoke-Expression",

		"'cmd /c start powershell -windowstyle hidden -c "+
			strings.Join(arg, " ")+
			"'",
	)
}

func CrossCmd(arg ...string) (*exec.Cmd, error) {

	c := exec.Command(
		"powershell",
		append([]string{"-c"}, arg...)...,
	)

	return c, nil

}

func GetDefaultTextEditor() string {
	return "notepad"
}

func ConfigFolder() string {

	configFolder, _ := os.LookupEnv("APPDATA")
	configFolder = filepath.Join(configFolder, "/jim")

	return configFolder
}

func AdminCmd(arg ...string) (*exec.Cmd, error) {

	if !isRunningAsAdmin() {
		runMeElevated()
		os.Exit(0)
	}

	return CrossCmd(arg...)
}
