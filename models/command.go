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

func (command *Command) AfterDelete(tx *gorm.DB) error {

	return tx.Model(&Session{}).
		Where("command_id = ?", command.ID).
		Unscoped().Delete(&Session{}).Error
}
