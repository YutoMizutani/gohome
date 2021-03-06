package repository

import (
	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/entity/primitive"
)

type TodoRepository struct {
	DataStore datastore.TodoDataStore
}

func (repository *TodoRepository) Add(entity *entity.Todo) error {
	return repository.DataStore.Create(entity)
}

func (repository *TodoRepository) GetAll() (*entity.TodoList, error) {
	return repository.DataStore.ReadAll()
}

func (repository *TodoRepository) Get(id *primitive.GormModelID) (*entity.Todo, error) {
	return repository.DataStore.Read(id)
}

func (repository *TodoRepository) Update(entity *entity.Todo) (*entity.Todo, error) {
	return repository.DataStore.Update(entity)
}

func (repository *TodoRepository) Delete(id *primitive.GormModelID) error {
	return repository.DataStore.Delete(id)
}
