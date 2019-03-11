package builder

import (
	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/data/repository"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
	"github.com/YutoMizutani/gohome/app/presenter/controller"
)

type TodoBuilder struct {
}

func (builder *TodoBuilder) Build() *controller.TodoController {
	todoDataStore := datastore.TodoDataStore{}
	err := todoDataStore.Migrate()
	if err != nil {
		panic(err)
	}
	return &controller.TodoController{
		UseCase: &usecase.TodoUseCase{
			Repository: &repository.TodoRepository{
				DataStore: todoDataStore,
			},
		},
	}
}
