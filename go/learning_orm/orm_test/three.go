package orm_test

import (
	"fmt"
	"learning_orm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := learning_orm.NewEngine("sqlite3", "test.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("create table User(Name text);").Exec()
	_, _ = s.Raw("create table User(Name text);").Exec()
	result, _ := s.Raw("insert into User(name) values (?), (?)", "Tom", "Jack").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
