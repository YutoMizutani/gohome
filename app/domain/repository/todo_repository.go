package repository

import "github.com/YutoMizutani/gohome/app/domain/entity"

type TodoRepository interface {
	Add(entity *entity.Todo) error
	GetAll() (*entity.TodoList, error)
	Get(id uint) (*entity.Todo, error)
	Update(entity *entity.Todo) (*entity.Todo, error)
	Delete(id uint) error
}
