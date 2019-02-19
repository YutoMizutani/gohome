package builder

import (
	"github.com/YutoMizutani/gohome/app/data/repository"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
	"github.com/YutoMizutani/gohome/app/presenter/controller"
)

type AnimalBuilder struct {
}

func (builder *AnimalBuilder) Build() *controller.AnimalController {
	return &controller.AnimalController{
		UseCase: &usecase.AnimalUseCase{
			Repository: &repository.AnimalRepository{},
		},
	}
}
