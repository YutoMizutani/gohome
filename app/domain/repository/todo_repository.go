package repository

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/entity/primitive"
)

type TodoRepository interface {
	Add(entity *entity.Todo) error
	GetAll() (*entity.TodoList, error)
	Get(id *primitive.GormModelID) (*entity.Todo, error)
	Update(entity *entity.Todo) (*entity.Todo, error)
	Delete(id *primitive.GormModelID) error
}
