package orm_test
// package main

import (
	"database/sql"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "test.db")
	defer func() { _= db.Close() }()
	_, _ = db.Exec("create table if not exists User(`Name` text);")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("insert into User(Name) values (?)", "Tom")
	_, err2 := tx.Exec("insert into User(Name) values (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}
}