package models

import (
	"time"

	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Value       string
	LastTouched time.Time
}
