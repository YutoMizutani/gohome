package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/repository"
)

type TodoUseCase struct {
	Repository repository.TodoRepository
}

func (usecase *TodoUseCase) Get() (*entity.TodoList, error) {
	entities, err := usecase.Repository.Get()
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (usecase *TodoUseCase) Add(entity *entity.Todo) error {
	err := usecase.Repository.Add(entity)
	if err != nil {
		return err
	}

	return nil
}
