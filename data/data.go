package data

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
)

var DB *sql.DB

func Start() {
	if DB != nil {
		panic(errors.New("DB already initialised"))
	}

	db, err := sql.Open("sqlite3", "./chinook.db")
	if err != nil {
		panic(err)
	}

	DB = db
	log.Print("Connected to database")
}

func Stop() {
	if DB == nil {
		panic(errors.New("DB not initialised"))
	}

	err := DB.Close()
	if err != nil {
		panic(err)
	}

	log.Print("Disconnected from database")
}

func Query(path string, args ...interface{}) sql.Rows {
	assertConnection()

	rows, err := DB.Query(readFile(path), args...)
	if err != nil {
		panic(err)
	}

	log.Print("Ran ", path, " with params ", args)

	return *rows
}

func Exec(path string, args ...interface{}) {
	assertConnection()

	_, err := DB.Exec(readFile(path), args...)
	if err != nil {
		panic(err)
	}

	log.Print("Ran ", path, " with params ", args)
}

func assertConnection() {
	if DB == nil {
		panic(errors.New("No connection to database"))
	}
}

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sql, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(sql)
}
