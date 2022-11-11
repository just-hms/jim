package models

import (
	"errors"
	"strings"

	"github.com/tidwall/buntdb"
)

type Command struct {
	Name  string
	Value string
}

func GetCommands(filter string, commands *[]Command) error {

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

func GetCommandByName(command *Command, name string) error {

	err := DB().View(func(tx *buntdb.Tx) error {
		var err error
		command.Value, err = tx.Get("command:" + name)
		command.Name = name // set the name to the key if found
		return err
	})

	return err
}

func (self *Command) Save() error {

	err := DB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("command:"+self.Name, self.Value, nil)
		return err
	})

	return err
}

func Clear() {
	DB().Update(func(tx *buntdb.Tx) error {
		tx.DeleteAll()
		return nil
	})
}

func (self *Command) Remove() error {

	err := DB().Update(func(tx *buntdb.Tx) error {

		// delete all sessions
		var delkeys []string

		tx.AscendKeys("session:"+self.Name+":*", func(k, v string) bool {
			delkeys = append(delkeys, k)
			return true // continue
		})

		for _, k := range delkeys {
			if _, err := tx.Delete(k); err != nil {
				return err
			}
		}

		_, err := tx.Delete("command:" + self.Name)
		return err
	})

	if err != nil {
		err = errors.New("error while deleting")
	}

	return err
}

func (self *Command) Rename(new_name string) error {

	err := DB().Update(func(tx *buntdb.Tx) error {

		var rnKeys []string

		// iterate over the sessions

		tx.AscendKeys("session:"+self.Name+":*", func(k, v string) bool {
			rnKeys = append(rnKeys, k)
			return true // continue
		})

		for _, k := range rnKeys {

			// delete the session
			val, delErr := tx.Delete(k)

			if delErr != nil {
				return delErr
			}

			// set the session value to the new key

			_, _, setErr := tx.Set(
				strings.ReplaceAll(k, self.Name, new_name),
				val, // the old session value
				nil,
			)

			if setErr != nil {
				return setErr
			}

		}

		// delete the old commands key and set the new one
		val, err := tx.Delete("command:" + self.Name)

		if err != nil {
			return err
		}

		_, _, setErr := tx.Set("command:"+new_name, val, nil)
		return setErr
	})

	return err
}
