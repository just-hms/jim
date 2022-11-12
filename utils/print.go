package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var Alertf = color.New(color.FgRed).PrintfFunc()
var Titlef = color.New(color.FgHiWhite, color.Bold).PrintfFunc()
var Commentf = color.New(color.FgHiBlack, color.Bold).PrintfFunc()
var Warningf = color.New().PrintfFunc()

func init() {

	if os.Getenv("testing") != "true" {
		return
	}

	var defaultf = func(format string, a ...interface{}) {
		fmt.Printf(format, a...)
	}

	Alertf = defaultf
	Titlef = defaultf
	Commentf = defaultf
	Warningf = defaultf
}
