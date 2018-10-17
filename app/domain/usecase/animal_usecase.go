package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/model"
	"github.com/YutoMizutani/gohome/app/domain/repository"
	"github.com/YutoMizutani/gohome/app/domain/translator"
)

type AnimalUseCase struct {
	Repository repository.AnimalRepository
	Translator translator.AnimalTranslator
}

func (usecase *AnimalUseCase) Fetch() (*model.AnimalModel, error) {
	entity, err := usecase.Repository.Fetch()
	if err != nil {
		return nil, err
	}

	animalModel := usecase.Translator.Translate(entity)
	return animalModel, nil
}
