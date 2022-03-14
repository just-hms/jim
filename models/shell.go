package models

import "gorm.io/gorm"

type Shell struct {
	gorm.Model
	Name      string `gorm:"uniqueIndex"`
	IsDefault bool
}
