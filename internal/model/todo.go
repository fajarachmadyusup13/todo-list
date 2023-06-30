package model

import "context"

type (
	Todo struct {
		ID          int64  `json:"id"`
		Description string `json:"description"`
	}

	TodoRepository interface {
		Create(ctx context.Context, todo *Todo) error
		FindByID(ctx context.Context, id int64) (res *Todo, err error)
		UpdateByID(ctx context.Context, todo *Todo) (res *Todo, err error)
		DeleteByID(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (ids []int64, err error)
	}

	TodoUsecase interface {
		Create(ctx context.Context, todo *Todo) error
		FindByID(ctx context.Context, id int64) (res *Todo, err error)
		UpdateByID(ctx context.Context, todo *Todo) (res *Todo, err error)
		DeleteByID(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (ids []int64, err error)
	}
)
