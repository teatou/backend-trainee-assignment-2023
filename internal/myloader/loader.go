package myloader

import (
	"database/sql"
	"fmt"
	"time"
)

type Loader interface {
	GetTransaction() (*sql.Tx, error)
	AddSegment(slug string) error
	RemoveSegment(slug string) error
	AddUserSegment(userId int, segmentSlug string) error
	RemoveUserSegment(userId int, segmentSlug string) error
	GetUserActiveSegments(userId int) (slugs []string, err error)
	AddUser(id int) error
	RemoveUser(id int) error
	GetUserIdPercent(percent int) (userIds []int, err error)
}

type Connector interface {
	Exec(query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func addSegment(connector Connector, slug string) (err error) {
	req := `insert into "segments"("slug") values($1)`
	_, err = connector.Exec(req, slug)
	return
}

func removeSegment(connector Connector, slug string) (err error) {
	req := `delete from "segments" where slug=$1`
	_, err = connector.Exec(req, slug)
	return
}

func addUser(connector Connector, id int) (err error) {
	req := `insert into "users"("id") values($1)`
	_, err = connector.Exec(req, id)
	return
}

func removeUser(connector Connector, id int) (err error) {
	req := `delete from "users" where id=$1`
	_, err = connector.Exec(req, id)
	return
}

func addUserSegment(connector Connector, userId int, segmentSlug string) (err error) {
	req := `insert into "user_segments"("user_id", "segment_slug",  "adding_time") values($1, $2, $3)`
	_, err = connector.Exec(req, userId, segmentSlug, time.Now())
	return
}

func removeUserSegment(connector Connector, userId int, segmentSlug string) (err error) {
	req := `update "user_segments" SET "removal_time" = $1 WHERE "user_id" = $2 and "segment_slug" = $3 and "removal_time" is NULL`
	_, err = connector.Exec(req, time.Now(), userId, segmentSlug)
	return
}

func getUserActiveSegments(connector Connector, userId int) (res []string, err error) {
	req := `SELECT "segment_slug" FROM "user_segments" WHERE "user_id" = $1 AND "removal_time" IS NULL`
	rows, err := connector.Query(req, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var slug string
		err = rows.Scan(&slug)
		if err != nil {
			return nil, fmt.Errorf("scaning rows: %w", err)
		}
		res = append(res, slug)
	}
	return
}

func getUserIdPercent(connector Connector, percent int) (userIds []int, err error) {
	req := `SELECT "id" FROM "users" TABLESAMPLE SYSTEM ($1)`
	rows, err := connector.Query(req, percent)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scaning rows: %w", err)
		}
		userIds = append(userIds, id)
	}
	return
}
