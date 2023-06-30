package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/todo-list/internal/model"
)

type toDoUsecase struct {
}

func NewToDoUsecase() model.TodoUsecase {
	return &toDoUsecase{}
}

func (t *toDoUsecase) Create(ctx context.Context, todo *model.Todo) error {
	return nil
}

func (t *toDoUsecase) FindByID(ctx context.Context, id int64) (res *model.Todo, err error) {
	return nil, nil
}

func (t *toDoUsecase) UpdateByID(ctx context.Context, todo *model.Todo) (res *model.Todo, err error) {
	return nil, nil
}

func (t *toDoUsecase) DeleteByID(ctx context.Context, id int64) error {
	return nil
}

func (t *toDoUsecase) FindAll(ctx context.Context) (ids []int64, err error) {
	return nil, nil
}
