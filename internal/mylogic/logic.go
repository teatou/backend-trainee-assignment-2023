package mylogic

import (
	"backend-trainee-assignment-2023/internal/myloader"
	"fmt"
)

type Logic struct {
	loader myloader.Loader
}

func NewLogic(ldr myloader.Loader) *Logic {
	return &Logic{
		loader: ldr,
	}
}

func (l *Logic) AddSegment(slug string) (err error) {
	err = l.loader.AddSegment(slug)
	if err != nil {
		err = fmt.Errorf("loader: %w", err)
	}
	return
}

func (l *Logic) RemoveSegment(slug string) (err error) {
	err = l.loader.RemoveSegment(slug)
	if err != nil {
		err = fmt.Errorf("loader: %w", err)
	}
	return
}

func (l *Logic) AddUser(id int) (err error) {
	err = l.loader.AddUser(id)
	if err != nil {
		err = fmt.Errorf("loader: %w", err)
	}
	return
}

func (l *Logic) RemoveUser(id int) (err error) {
	err = l.loader.RemoveUser(id)
	if err != nil {
		err = fmt.Errorf("loader: %w", err)
	}
	return
}

func (l *Logic) UpdateUserSegments(userId int, addingSlugs []string, removalSlugs []string) error {
	tx, err := l.loader.GetTransaction()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	txLoader := myloader.NewTxLoader(tx)
	for i := range addingSlugs {
		if err = txLoader.AddUserSegment(userId, addingSlugs[i]); err != nil {
			tx.Rollback()
			return fmt.Errorf("adding user segement: %w", err)
		}
	}
	for i := range removalSlugs {
		if err = txLoader.RemoveUserSegment(userId, removalSlugs[i]); err != nil {
			tx.Rollback()
			return fmt.Errorf("removing user segement: %w", err)
		}
	}
	tx.Commit()
	return nil
}

func (l *Logic) GetUserActiveSegments(userId int) ([]string, error) {
	res, err := l.loader.GetUserActiveSegments(userId)
	if err != nil {
		return nil, fmt.Errorf("loader: %w", err)
	}
	return res, nil
}

func (l *Logic) AddSegmentForUsersPercent(slug string, percent int) error {
	tx, err := l.loader.GetTransaction()
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}
	txLoader := myloader.NewTxLoader(tx)
	userIds, err := txLoader.GetUserIdPercent(percent)
	if err != nil {
		return fmt.Errorf("getting user ids: %w", err)
	}
	for _, id := range userIds {
		if err = txLoader.AddUserSegment(id, slug); err != nil {
			tx.Rollback()
			return fmt.Errorf("adding segment for user: %w", err)
		}
	}
	tx.Commit()
	return nil
}
