package orm_test_test

import (
	"errors"
	"learning_orm"
	"learning_orm/session"
	"reflect"
	"testing"
)

func OpenDB(t *testing.T) *learning_orm.Engine {
	t.Helper()
	engine, err := learning_orm.NewEngine("sqlite3", "test.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

type User struct {
	Name string `learning_orm:"PRIMARY KEY"`
	Age int
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T){
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(
		func(s *session.Session)(result interface{}, err error){
			_ = s.Model(&User{}).CreateTable()
			_, err = s.Insert(&User{"Tom", 18})
			return nil, errors.New("error")
		})

	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}


func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(
		func (s *session.Session) (result interface{}, err error) {
			_ = s.Model(&User{}).CreateTable()
			_, err = s.Insert(&User{"Tom", 18})
			return
		})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("create table User(Name text primary key, xxx integer);").Exec()
	_, _ = s.Raw("insert into User(Name) values (?), (?)", "Tom", "Sam").Exec()
	engine.Migrate(&User{})

	rows, _ := s.Raw("select * from User").QueryRows()
	columns, _ := rows.Columns()
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", columns)
	}
}