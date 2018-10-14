package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/model"
	"github.com/YutoMizutani/gohome/app/domain/repository"
	"github.com/YutoMizutani/gohome/app/domain/translator"
)

type AnimalUsecase struct {
	Repository repository.AnimalRepository
	Translator translator.AnimalTranslator
}

func (usecase *AnimalUsecase) Fetch() (model model.AnimalModel, err error) {
	entity, err := usecase.Repository.Fetch()
	model = usecase.Translator.Translate(&entity)
	return
}
