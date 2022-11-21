package actions

import (
	"fmt"
	"jim/internal/constants"
	"jim/pkg/io"
	"jim/pkg/rainbow"
	"log"
	"net/http"
	"os"
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

		// if it is not the last upgrgade jim

		exe_folder := io.ExecutableFolder()
		exe_path := io.Executable()

		update_link := "https://github.com/just-hms/jim/releases/latest/download/jim-" + runtime.GOOS + "-amd64.tar.gz"
		tmp_dir := os.TempDir() + "/jim.tar.gz"

		command := "curl -L " + update_link + " > " + tmp_dir + " ; " +
			"tar -xvf " + tmp_dir + " -C " + exe_folder

		if runtime.GOOS == "windows" {
			command = "curl " + update_link + " -O " + tmp_dir + " ; " +
				"sudo tar -xvf " + tmp_dir + " -C " + exe_folder
		}

		c, err := io.AdminCmd(command)

		if err != nil {
			rainbow.Alertf("%s\n", err.Error())
			return
		}

		// rename the executable so it doesn't brake

		os.Rename(exe_path, exe_path+".old")

		// run the upgrade

		if err := c.Run(); err != nil {
			rainbow.Alertf("%s\n", err.Error())
			os.Rename(exe_path+".old", exe_path)
			return
		}
	},
	Description:     "upgrade jim",
	HelpDescription: " Upgrade jim using this syntax\n\n     jim --upgrade\n\n If you have installed the last version this action will do nothing",

	ArgumentsCheck: func(args []string) bool {
		return len(args) == 0
	},
}
