package repository

import "github.com/YutoMizutani/gohome/app/domain/entity"

type TodoRepository interface {
	Get() (*entity.TodoList, error)
	Add(entity *entity.Todo) error
}
