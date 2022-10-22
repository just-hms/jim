package models

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB = nil

func DB() (db *gorm.DB) {

	if database != nil {
		return database
	}

	dbFolder := ""

	if runtime.GOOS == "windows" {
		dbFolder, _ = os.LookupEnv("APPDATA")
		dbFolder = filepath.Join(dbFolder, "/jim")
	} else {
		dbFolder, _ = os.LookupEnv("HOME")
		dbFolder = filepath.Join(dbFolder, "/.local/share/jim")
	}

	os.MkdirAll(dbFolder, os.ModePerm)

	dbName := filepath.Join(dbFolder, "/jim.db")

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // speed up
		Logger:                                   logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	database = db
	return database
}

func Build() {

	DB().AutoMigrate(
		&Command{},
	)
}
