package myloader

import (
	"database/sql"
)

type TxLoader struct {
	tx *sql.Tx
}

func NewTxLoader(tx *sql.Tx) Loader {
	return &TxLoader{
		tx: tx,
	}
}

func (l *TxLoader) GetTransaction() (*sql.Tx, error) {
	//TODO or make error?
	l.tx.Commit()
	return l.tx, nil
}

func (l *TxLoader) AddSegment(slug string) error {
	return addSegment(l.tx, slug)
}

func (l *TxLoader) RemoveSegment(slug string) error {
	return removeSegment(l.tx, slug)
}

func (l *TxLoader) AddUser(id int) error {
	return addUser(l.tx, id)
}

func (l *TxLoader) RemoveUser(id int) error {
	return removeUser(l.tx, id)
}

func (l *TxLoader) AddUserSegment(userId int, segmentSlug string) error {
	return addUserSegment(l.tx, userId, segmentSlug)
}

func (l *TxLoader) RemoveUserSegment(userId int, segmentSlug string) error {
	return removeUserSegment(l.tx, userId, segmentSlug)
}

func (l *TxLoader) GetUserActiveSegments(userId int) ([]string, error) {
	return getUserActiveSegments(l.tx, userId)
}

func (l *TxLoader) GetUserIdPercent(percent int) ([]int, error) {
	return getUserIdPercent(l.tx, percent)
}
