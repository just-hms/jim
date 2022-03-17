package main

import (
	"errors"
	"fmt"
	"jim/models"
	"jim/utils"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

func run(command models.Command) {
	command.LastTouched = time.Now()
	models.DB().Save(&command)

	var c *exec.Cmd

	if runtime.GOOS == "windows" {
		c = exec.Command("powershell", "-c", command.Value)
	} else {

		shell, err := os.LookupEnv("$SHELL")

		if !err {
			Alertf("no shell found!!!")
			return
		}

		c = exec.Command(shell, "-c", command.Value)
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {

		fmt.Println(err.Error())

		if exitError, ok := err.(*exec.ExitError); ok {

			exitCode := exitError.ExitCode()

			// TODO exit doesn't work

			if exitCode == 1 {
				Alertf("exit 1 test")
				go func() { os.Exit(0) }()
				func() {}()
			}
		}
	}
}

func find_command_by_name(name string, command *models.Command) error {

	if err := models.DB().Where("name = ?", name).First(command).Error; err == nil {
		return nil
	}

	// if no result was found try with similiar names

	commands := []models.Command{}
	models.DB().Find(&commands)

	if len(commands) == 0 {
		fmt.Println("...")
		return errors.New("commands empty")
	}

	max_lev_rateo := 0.0

	for i, item := range commands {

		if len(item.Name) == 0 {
			continue
		}

		item_lev :=
			float64((len(item.Name) - utils.Levenshtein(item.Name, name))) /
				float64(len(item.Name))

		if item_lev > max_lev_rateo {
			max_lev_rateo = item_lev
			*command = commands[i]
		}
	}

	if max_lev_rateo == 0 {
		fmt.Println("...")
		return errors.New("nothing similiar found")

	}

	if max_lev_rateo >= 0.80 {
		return nil
	}

	Warningf("did you mean %s? Type y or N\n", command.Name)

	if utils.ReadChar() == 'y' {
		return nil
	}

	return errors.New("not found")

}

type Action struct {
	Value        func([]string)
	Description  string
	ArgumentsLen int
}

var actions = map[string]*Action{
	"--ls": {
		Value: func(args []string) {

			commands := []models.Command{}
			models.DB().Table("commands").Scan(&commands)

			if len(commands) == 0 {
				fmt.Println("...")
				return
			}

			Titlef("%-10s%-30s%-1s\n", "ID", "Name", "Comand")

			for _, v := range commands {
				fmt.Printf("%-10d%-30s%-1s\n", v.ID, v.Name, v.Value)
			}

		},
		Description:  "list of all commands",
		ArgumentsLen: 0,
	},
	"--add": {
		Value: func(args []string) {

			to_search := models.Command{}

			if err := models.DB().Where("name = ?", args[0]).First(&to_search).Error; err == nil {
				Alertf("a command named %s already exists!!!\n", args[0])
				return
			}

			args[1] = strings.Replace(args[1], "$", utils.CurrentFolder(), -1)

			command := models.Command{
				Name:        args[0],
				Value:       args[1],
				LastTouched: time.Now(),
			}

			models.DB().Create(&command)

		},
		Description:  "add a command",
		ArgumentsLen: 2,
	},
	"--mod": {
		Value: func(args []string) {
			args[1] = strings.Replace(args[1], "$", utils.CurrentFolder(), -1)

			command := models.Command{}

			if err := find_command_by_name(args[0], &command); err != nil {
				return
			}

			command.Value = args[1]
			models.DB().Save(&command)
		},
		Description:  "modify a specified command",
		ArgumentsLen: 2,
	},
	"--modi": {
		Value: func(args []string) {
			args[1] = strings.Replace(args[1], "$", utils.CurrentFolder(), -1)

			command := models.Command{}

			if err := models.DB().Where("id = ?", args[0]).First(&command).Error; err != nil {
				Alertf("id not found")
				return
			}

			command.Value = args[1]
			models.DB().Save(&command)
		},
		Description:  "modify a command ny id",
		ArgumentsLen: 2,
	},
	"--rm": {
		Value: func(args []string) {
			command := models.Command{}

			if err := find_command_by_name(args[0], &command); err != nil {
				return
			}

			models.DB().Unscoped().Delete(&command)
		},
		Description:  "remove a command",
		ArgumentsLen: 1,
	},
	"--rmi": {
		Value: func(args []string) {
			command := models.Command{}

			if err := models.DB().Where("id = ?", args[0]).First(&command).Error; err != nil {
				Alertf("id not found")
				return
			}

			models.DB().Unscoped().Delete(&command)
		},
		Description:  "remove a command by id",
		ArgumentsLen: 1,
	},
	"--cls": {
		Value: func(args []string) {

			Alertf("clear all commands is not reversible, are you sure? Type y or N\n")

			if utils.ReadChar() != 'y' {
				return
			}

			models.DB().Unscoped().Where("1=1").Delete(&models.Command{})
		},
		Description:  "clear all commands",
		ArgumentsLen: 0,
	},
	"--run": {
		Value: func(args []string) {
			command := models.Command{}

			if err := find_command_by_name(args[0], &command); err != nil {
				return
			}

			run(command)

		},
		Description:  "run a command",
		ArgumentsLen: 1,
	},
	"--rn": {
		Value: func(args []string) {
			command := models.Command{}

			if err := find_command_by_name(args[0], &command); err != nil {
				return
			}

			to_rename := models.Command{}

			if err := models.DB().Where("name = ?", args[1]).First(&to_rename).Error; err == nil {
				Alertf("a command named %s already exists!!!\n", args[1])
				return
			}

			models.DB().Model(&command).Update("name", args[1])
		},
		Description:  "rename a command",
		ArgumentsLen: 2,
	},
	"--cache": {
		Value: func(args []string) {
			command := models.Command{}

			if err := models.DB().Order("last_touched").First(&command).Error; err != nil {
				fmt.Println("...")
				return
			}

			run(command)
		},
		Description:  "run the last touched command",
		ArgumentsLen: 0,
	},
}

var Alertf = color.New(color.FgRed).PrintfFunc()
var Titlef = color.New(color.FgHiWhite, color.Bold).PrintfFunc()
var Warningf = color.New().PrintfFunc()

func main() {

	actions["--help"] = &Action{
		Value: func(args []string) {

			Titlef("%-20s%-20s\n", "Command", "Descripion")

			keys := make([]string, 0, len(actions))
			for k := range actions {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			for _, key := range keys {
				fmt.Printf("%-20s%-20s\n", key, actions[key].Description)
			}
		},
		Description:  "list of all possible actions and their descriprion",
		ArgumentsLen: 0,
	}

	models.Build()

	if len(os.Args) <= 1 {
		return
	}

	command := os.Args[1]
	var args []string

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	action := actions[command]

	if action == nil {

		// TODO : somthing like append([]string{command}, args...),

		actions["--run"].Value([]string{command})
		return
	}

	if action.ArgumentsLen != len(args) {
		Alertf("wrong format!!!")
		return
	}

	action.Value(args)
}
