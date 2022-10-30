package models

import (
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"` // kek
	Value    string // code ...
	Sessions []Session
}
