package learning_orm

import (
	"database/sql"
	"fmt"
	"learning_orm/dialect"
	"learning_orm/log"
	"learning_orm/session"
	"strings"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

type TxFunc func(* session.Session) (interface{}, error)

// NewEngine create an new engine
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

// Close enging
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}

	log.Info("Close database success")
}

// NewSession create session
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

func (engine *Engine) Transaction(f TxFunc)(result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p:= recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()

	return f(s)
}

func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

func (engine *Engine) Migrate(value interface{}) error {
	_, err := engine.Transaction(
		func(s *session.Session) (result interface{}, err error) {
			if !s.Model(value).HasTable() {
				log.Info("table %s doesn't exist", s.RefTable().Name)
				return nil, s.CreateTable()
			}
			table := s.RefTable()
			rows, _ := s.Raw(fmt.Sprintf("select * from %s limit 1", table.Name)).QueryRows()
			columns, _ := rows.Columns()
			addCols := difference(table.FieldNames, columns)
			delCols := difference(columns, table.FieldNames)
			log.Infof("added cols %v, deleted cols %v", addCols, delCols)

			for _, col := range addCols {
				f := table.GetField(col)
				sqlStr := fmt.Sprintf("alter table %s add column %s %s;", table.Name, f.Name, f.Type)
				if _, err = s.Raw(sqlStr).Exec(); err != nil {
					return
				}
			}

			if len(delCols) == 0 {
				return
			}
			tmp := "tmp_" + table.Name
			fieldStr := strings.Join(table.FieldNames, ", ")
			s.Raw(fmt.Sprintf("create table %s as select %s from %s;", tmp, fieldStr, table.Name))
			s.Raw(fmt.Sprintf("drop table %s;", table.Name))
			s.Raw(fmt.Sprintf("alter table %s rename to %s", tmp, table.Name))
			_, err = s.Exec()
			return
		})
	return err
}