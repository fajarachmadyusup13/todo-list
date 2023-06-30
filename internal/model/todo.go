package model

import "context"

type (
	ToDo struct {
		ID          int64  `json:"id"`
		Description string `json:"description"`
	}

	ToDoRepository interface {
		Create(ctx context.Context, todo *ToDo) error
		FindByID(ctx context.Context, id int64) (res *ToDo, err error)
		UpdateByID(ctx context.Context, todo *ToDo) (res *ToDo, err error)
		DeleteByID(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (ids []int64, err error)
	}

	ToDoUsecase interface {
		Create(ctx context.Context, todo *ToDo) error
		FindByID(ctx context.Context, id int64) (res *ToDo, err error)
		UpdateByID(ctx context.Context, todo *ToDo) (res *ToDo, err error)
		DeleteByID(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (ids []int64, err error)
	}
)
