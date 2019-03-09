package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/entity/primitive"
)

type TodoUseCase interface {
	GetAll() (*entity.TodoList, error)
	Get(id *primitive.GormModelID) (*entity.Todo, error)
	Add(entity *entity.Todo) error
	Update(entity *entity.Todo) (*entity.Todo, error)
	UpdateDoneState(id *primitive.GormModelID, isDone bool) (*entity.Todo, error)
	Delete(id *primitive.GormModelID) error
}
