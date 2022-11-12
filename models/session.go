package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/buntdb"
)

type Session struct {
	Start   time.Time
	Elapsed time.Duration
	Command string
}

func GetSessions(filter string, sessions *[]Session) error {

	return DB().View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("sessions", func(key, value string) bool {

			session := Session{}

			if err := json.Unmarshal([]byte(value), &session); err != nil {
				return false
			}

			if filter != "" && !strings.Contains(session.Command, filter) {
				return true
			}

			*sessions = append(*sessions, session)
			return true // continue iteration
		})

		return err
	})
}

func (self *Session) Save() error {

	err := DB().Update(func(tx *buntdb.Tx) error {

		var (
			b   []byte
			err error
		)

		b, err = json.Marshal(self)

		if err != nil {
			return err
		}

		fmt.Print(string(b))

		_, _, err = tx.Set(
			"session:"+self.Command+":"+strconv.FormatInt(self.Start.Unix(), 10),
			string(b),
			nil,
		)

		return err
	})

	return err
}
