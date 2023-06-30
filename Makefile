SHELL:=/bin/bash

internal/model/mock/mock_todo_repository.go:
	mockgen -destination=internal/model/mock/mock_todo_repository.go -package=mock github.com/fajarachmadyusup13/todo-list/internal/model TodoRepository

mockgen: internal/model/mock/mock_todo_repository.go

clean:
	rm -v internal/model/mock/mock_*_repository.go
	rm -v client/mock/mock*.go