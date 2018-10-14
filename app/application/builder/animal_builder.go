package builder

import (
	"gohome/app/data/repository"
	"gohome/app/domain/translator"
	"gohome/app/domain/usecase"
	"gohome/app/presenter/controller"
)

type AnimalBuilder struct {
}

func (builder *AnimalBuilder) Build() *controller.AnimalController {
	return &controller.AnimalController{
		Usecase: usecase.AnimalUsecase{
			Repository: &repository.AnimalRepository{},
			Translator: translator.AnimalTranslator{},
		},
	}
}
