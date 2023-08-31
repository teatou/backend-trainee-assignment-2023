package myloader

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PgLoader struct {
	db *sql.DB
}

func NewPgLoader(host string, port int, user, password, dbName string) (Loader, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, fmt.Errorf("database connetction: %w", err)
	}
	return &PgLoader{
		db: db,
	}, nil
}

func (l *PgLoader) GetTransaction() (*sql.Tx, error) {
	return l.db.Begin()
}

func (l *PgLoader) AddSegment(slug string) error {
	return addSegment(l.db, slug)
}

func (l *PgLoader) RemoveSegment(slug string) error {
	return removeSegment(l.db, slug)
}

func (l *PgLoader) AddUser(id int) error {
	return addUser(l.db, id)
}

func (l *PgLoader) RemoveUser(id int) error {
	return removeUser(l.db, id)
}

func (l *PgLoader) AddUserSegment(userId int, segmentSlug string) error {
	return addUserSegment(l.db, userId, segmentSlug)
}

func (l *PgLoader) RemoveUserSegment(userId int, segmentSlug string) error {
	return removeUserSegment(l.db, userId, segmentSlug)
}

func (l *PgLoader) GetUserActiveSegments(userId int) ([]string, error) {
	return getUserActiveSegments(l.db, userId)
}

func (l *PgLoader) GetUserIdPercent(percent int) ([]int, error) {
	return getUserIdPercent(l.db, percent)
}
