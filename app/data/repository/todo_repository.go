package repository

import (
	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/domain/entity"
)

type TodoRepository struct {
	DataStore datastore.TodoDataStore
}

func (repository *TodoRepository) Get() (*entity.TodoList, error) {
	return repository.DataStore.FindAll()
}

func (repository *TodoRepository) Add(entity *entity.Todo) error {
	return repository.DataStore.Create(entity)
}
