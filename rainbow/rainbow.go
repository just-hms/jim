package rainbow

import (
	"fmt"

	"github.com/fatih/color"
)

var Alertf = color.New(color.FgRed).PrintfFunc()
var Titlef = color.New(color.FgHiWhite, color.Bold).PrintfFunc()
var Commentf = color.New(color.FgHiBlack, color.Bold).PrintfFunc()
var Warningf = color.New().PrintfFunc()

// change the colored apis back to standard fmt
func Blank() {
	var defaultf = func(format string, a ...interface{}) {
		fmt.Printf(format, a...)
	}

	Alertf = defaultf
	Titlef = defaultf
	Commentf = defaultf
	Warningf = defaultf
}
