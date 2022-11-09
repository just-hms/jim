package models

import (
	"strings"

	"github.com/tidwall/buntdb"
)

type Command struct {
	Name  string
	Value string
}

func ListCommands(filter string, commands *[]Command) error {

	return DB().View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("commands", func(key, value string) bool {

			command := Command{
				Name:  strings.Split(key, "command:")[1],
				Value: value,
			}

			if filter != "" && !strings.Contains(command.Name, filter) {
				return true
			}

			*commands = append(*commands, command)
			return true // continue iteration
		})
		return err
	})
}
