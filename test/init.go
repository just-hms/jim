package test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func init() {
	testing.Init()
}

func getStdout(callback func()) string {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	callback()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return strings.TrimSpace(string(out))
}
