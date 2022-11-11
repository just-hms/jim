package models

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tidwall/buntdb"
)

var database *buntdb.DB = nil

func init() {
	DB().CreateIndex("commands", "command:*", buntdb.IndexString)
	DB().CreateIndex("sessions", "session:*:*", buntdb.IndexString)
}

func DB() (db *buntdb.DB) {

	if database != nil {
		return database
	}

	if os.Getenv("testing") == "true" {

		if db, err := buntdb.Open(":memory:"); err != nil {
			panic("failed to connect database")
		} else {
			database = db
		}

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
	dbName := filepath.Join(dbFolder, "/jim.kv.db")
	db, err := buntdb.Open(dbName)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	database = db
	return database
}
