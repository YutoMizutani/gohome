package usecase

import (
	"gohome/app/domain/model"
	"gohome/app/domain/repository"
	"gohome/app/domain/translator"
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
