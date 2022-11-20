package actions

import (
	"fmt"
	"jim/internal/constants"
	"jim/pkg/rainbow"
	"jim/pkg/utils"
	"log"
	"net/http"
	"os"
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
			fmt.Println("No applicable update found.")
			return
		}

		// rename the executable so it doesn't brake
		os.Rename(utils.Executable(), utils.Executable()+".old")

		// otherwise upgrade

		fmt.Println(utils.ExecutableFolder())

		update_link := "https://github.com/just-hms/jim/releases/latest/download/jim-" + runtime.GOOS + "-amd64.tar.gz"
		tmp_dir := os.TempDir() + "/jim.tar.gz"

		var c *exec.Cmd

		if runtime.GOOS == "windows" {
			c, err = utils.CrossCmd(
				"curl " + update_link + " -O " + tmp_dir + " ; " +
					"tar -xvf " + tmp_dir + " -C " + utils.ExecutableFolder(),
			)
		} else {
			c, err = utils.CrossCmd(
				"curl -L " + update_link + " > " + tmp_dir + " ; " +
					"tar -xvf " + tmp_dir + " -C " + utils.ExecutableFolder(),
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
	},
	Description:     "upgrade jim",
	HelpDescription: " Upgrade jim using this syntax\n\n     jim --upgrade\n\n If you have installed the last version this action will do nothing",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
}
