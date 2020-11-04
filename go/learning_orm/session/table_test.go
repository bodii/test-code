package session

import (
	"learning_orm"
	"testing"
)

type User struct {
	Name string `learning_orm:"PRIMARY KEY"`
	Age  int
}

func NewSession() *Session {
	engine, _ := learning_orm.NewEngine("sqlite3", "test.db")
	defer engine.Close()
	return engine.NewSession()
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()

	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
