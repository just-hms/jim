package actions

import (
	"jim/internal/constants"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

var Upgrade = &Action{
	Value: func(args []string) {

		// check if the version of jim is the last

		resp, err := http.Get("https://github.com/just-hms/jim/releases/latest")
		if err != nil {
			log.Fatalln(err)
		}

		path := strings.Split(resp.Request.URL.Path, "/")

		last_version := path[len(path)-1]

		if last_version == constants.Version {
			return
		}

		// otherwise upgrade

		utils.ExecutableFolder()

		var c *exec.Cmd

		if runtime.GOOS == "windows" {

		} else if runtime.GOOS == "darwin" {

		} else {
			c, err = utils.CrossCmd(
				"\"curl -L https://github.com/just-hms/jim/releases/latest/download/jim-linux-amd64.tar.gz > /tmp/jim.tar.gz ;",
				"sudo tar -xvf /tmp/jim.tar.gz -C "+utils.ExecutableFolder()+"\"",
			)
		}

		if err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		if err := c.Run(); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		// call jim --version
		Version.Value([]string{})
	},
	Description:     "upgrade jim",
	HelpDescription: " Upgrade jim using this syntax\n\n     jim --upgrade\n\n If you have installed the last version this action will do nothing",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
}
