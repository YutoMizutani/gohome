package usecase

import "github.com/YutoMizutani/gohome/app/domain/entity"

type TodoUseCase interface {
	GetAll() (*entity.TodoList, error)
	Get(id uint) (*entity.Todo, error)
	Add(entity *entity.Todo) error
	Update(entity *entity.Todo) (*entity.Todo, error)
	UpdateDoneState(id uint, isDone bool) (*entity.Todo, error)
	Delete(id uint) error
}
