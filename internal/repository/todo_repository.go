package repository

import (
	"context"
	"errors"

	"github.com/fajarachmadyusup13/todo-list/internal/model"
	"github.com/fajarachmadyusup13/todo-list/internal/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) model.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

// Create :nodoc:
func (t *todoRepository) Create(ctx context.Context, todo *model.Todo) error {
	if todo.ID != 0 {
		return errors.New("invalid data")
	}

	todo.ID = utils.GenerateID()

	tx := t.db.WithContext(ctx).Begin()
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.DumpIncomingContext(ctx),
		"todo":    utils.Dump(todo),
	})

	err := tx.Create(&todo).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		logger.WithField("tx", utils.Dump(tx)).Error(err)
		return err
	}

	return nil
}

// FindByID :nodoc:
func (t *todoRepository) FindByID(ctx context.Context, id int64) (res *model.Todo, err error) {
	return nil, nil
}

func (t *todoRepository) UpdateByID(ctx context.Context, todo *model.Todo) (res *model.Todo, err error) {
	return nil, nil
}

func (t *todoRepository) DeleteByID(ctx context.Context, id int64) error {
	return nil
}

func (t *todoRepository) FindAll(ctx context.Context) (ids []int64, err error) {
	return nil, nil
}
