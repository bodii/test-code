package learning_orm

import (
	"database/sql"
	"learning_orm/log"
	"learning_orm/session"
)

type Engine struct {
	db *sql.DB
}

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
	e = &Engine{db: db}
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
	return session.New(engine.db)
}
