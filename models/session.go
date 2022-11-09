package models

import (
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

func ListSessions(filter string, sessions *[]Session) error {

	return DB().View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("sessions", func(key, value string) bool {

			key_data := strings.Split(strings.TrimPrefix(key, "session:"), ":")

			start_unix, _ := strconv.ParseInt(key_data[1], 10, 64)
			elapsed, _ := strconv.ParseInt(value, 10, 64)

			session := Session{
				Start:   time.Unix(start_unix, 0),
				Command: key_data[0],
				Elapsed: time.Duration(elapsed) * time.Millisecond,
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
