package usecase

import "github.com/YutoMizutani/gohome/app/domain/entity"

type TodoUseCase interface {
	Get() (*entity.TodoList, error)
	Add(entity *entity.Todo) error
}
