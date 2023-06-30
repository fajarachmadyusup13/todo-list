package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarachmadyusup13/todo-list/internal/model"
	"github.com/fajarachmadyusup13/todo-list/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newDummyToDo() *model.Todo {
	return &model.Todo{
		Description: "desc test",
	}
}

func initializeToDoRepositoryWithMock(dbGorm *gorm.DB) (tr model.TodoRepository) {
	initializeTest()

	tr = NewToDoRepository(dbGorm)

	return
}

func TestTodoRepo_Create(t *testing.T) {
	ctx := context.TODO()

	t.Run("create success", func(t *testing.T) {
		todo := newDummyToDo()

		db, dbmock := initializeCockroachMockConn()
		tr := initializeToDoRepositoryWithMock(db)

		dbmock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(todo.ID)
		dbmock.ExpectQuery("INSERT INTO \"todos\"").
			WithArgs(todo.Description, sqlmock.AnyArg()).WillReturnRows(queryResult)
		dbmock.ExpectCommit()

		err := tr.Create(ctx, todo)

		assert.NoError(t, err)
	})

	t.Run("create failed invalid data", func(t *testing.T) {
		todo := newDummyToDo()
		todo.ID = utils.GenerateID()

		db, _ := initializeCockroachMockConn()
		tr := initializeToDoRepositoryWithMock(db)

		err := tr.Create(ctx, todo)
		assert.Equal(t, errors.New("invalid data"), err)
	})

	t.Run("create failed begin failure", func(t *testing.T) {
		todo := newDummyToDo()

		db, dbMock := initializeCockroachMockConn()
		tr := initializeToDoRepositoryWithMock(db)

		dbMock.ExpectBegin().WillReturnError(gorm.ErrInvalidTransaction)
		dbMock.ExpectQuery("INSERT INTO \"todos\"").WillReturnError(gorm.ErrInvalidTransaction)
		dbMock.ExpectRollback()

		err := tr.Create(ctx, todo)

		assert.Equal(t, gorm.ErrInvalidTransaction, err)
	})

	t.Run("create failed unaddresable", func(t *testing.T) {
		todo := newDummyToDo()

		db, dbMock := initializeCockroachMockConn()
		tr := initializeToDoRepositoryWithMock(db)

		dbMock.ExpectBegin()
		dbMock.ExpectQuery("INSERT INTO \"todos\"").WillReturnError(gorm.ErrInvalidDB)
		dbMock.ExpectRollback()

		err := tr.Create(ctx, todo)

		assert.Equal(t, gorm.ErrInvalidDB, err)
	})

	t.Run("create failed commit failure", func(t *testing.T) {
		todo := newDummyToDo()

		db, dbMock := initializeCockroachMockConn()
		tr := initializeToDoRepositoryWithMock(db)

		dbMock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(todo.ID)
		dbMock.ExpectQuery("INSERT INTO \"todos\"").
			WithArgs(todo.Description, sqlmock.AnyArg()).WillReturnRows(queryResult)
		dbMock.ExpectCommit().WillReturnError(gorm.ErrInvalidTransaction)

		err := tr.Create(ctx, todo)

		assert.Equal(t, gorm.ErrInvalidTransaction, err)
	})
}
