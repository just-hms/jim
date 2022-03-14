package models

import (
	"jim/utils"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB = nil

func DB() (db *gorm.DB) {

	if database != nil {
		return database
	}

	dbName := filepath.Join(utils.ExecutableFolder(), "jim.db")

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // speed up
		Logger:                                   logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}

	database = db
	return database
}

func Build() {

	DB().AutoMigrate(
		&Command{},
		&Shell{},
	)

	// TODO : edit this
	var shell Shell

	if err := DB().First(&shell).Error; err == nil {
		return
	}

	shell = Shell{
		Name:      "powershell",
		IsDefault: true,
	}

	DB().Create(&shell)
}
