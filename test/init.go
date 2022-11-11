package test

import (
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	testing.Init()
	os.Setenv("testing", "true")
}

func getStdout(callback func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	callback()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
