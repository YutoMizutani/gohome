package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/entity/primitive"
	"github.com/YutoMizutani/gohome/app/domain/repository"
)

type TodoUseCase struct {
	Repository repository.TodoRepository
}

func (usecase *TodoUseCase) GetAll() (*entity.TodoList, error) {
	entities, err := usecase.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (usecase *TodoUseCase) Get(id *primitive.GormModelID) (*entity.Todo, error) {
	entity, err := usecase.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (usecase *TodoUseCase) Add(entity *entity.Todo) error {
	err := usecase.Repository.Add(entity)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TodoUseCase) Update(entity *entity.Todo) (*entity.Todo, error) {
	entity, err := usecase.Repository.Update(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (usecase *TodoUseCase) UpdateDoneState(id *primitive.GormModelID, isDone bool) (*entity.Todo, error) {
	var entity *entity.Todo
	entity, err := usecase.Repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.IsDone = isDone
	entity, err = usecase.Repository.Update(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (usecase *TodoUseCase) Delete(id *primitive.GormModelID) error {
	err := usecase.Repository.Delete(id)
	return err
}
