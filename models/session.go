package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Elapsed time.Duration

	CommandID uint
	Command   Command
}
