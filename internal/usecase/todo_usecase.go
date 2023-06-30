package usecase

import (
	"CRUDWithCockroach/internal/model"
	"context"
)

type toDoUsecase struct {
}

func NewToDoUsecase() model.ToDoUsecase {
	return &toDoUsecase{}
}

func (t *toDoUsecase) Create(ctx context.Context, todo *model.ToDo) error {
	return nil
}

func (t *toDoUsecase) FindByID(ctx context.Context, id int64) (res *model.ToDo, err error) {
	return nil, nil
}

func (t *toDoUsecase) UpdateByID(ctx context.Context, todo *model.ToDo) (res *model.ToDo, err error) {
	return nil, nil
}

func (t *toDoUsecase) DeleteByID(ctx context.Context, id int64) error {
	return nil
}

func (t *toDoUsecase) FindAll(ctx context.Context) (ids []int64, err error) {
	return nil, nil
}
