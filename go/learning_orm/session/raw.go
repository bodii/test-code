package session

import (
	"database/sql"
	"strings"
)

// Session type struct
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

// New create Session instance
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear reset Session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB return db object
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw create sql content
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values)
	return s
}
