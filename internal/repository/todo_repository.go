package repository

import (
	"CRUDWithCockroach/internal/model"
	"context"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) model.ToDoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) Create(ctx context.Context, todo *model.ToDo) error {
	return nil
}

func (t *todoRepository) FindByID(ctx context.Context, id int64) (res *model.ToDo, err error) {
	return nil, nil
}

func (t *todoRepository) UpdateByID(ctx context.Context, todo *model.ToDo) (res *model.ToDo, err error) {
	return nil, nil
}

func (t *todoRepository) DeleteByID(ctx context.Context, id int64) error {
	return nil
}

func (t *todoRepository) FindAll(ctx context.Context) (ids []int64, err error) {
	return nil, nil
}
