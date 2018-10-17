package usecase_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/data/entity"
	"github.com/YutoMizutani/gohome/app/domain/translator"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
)

type animalRepositoryMock struct {
}

func (repository *animalRepositoryMock) Fetch() (*entity.AnimalEntity, error) {
	animalEntity := &entity.AnimalEntity{}
	animalEntity.Name = "Cat"
	return animalEntity, nil
}

func TestAnimalUseCaseFetch(t *testing.T) {
	mock := new(animalRepositoryMock)

	animalUseCase := usecase.AnimalUseCase{}
	animalUseCase.Repository = mock
	animalUseCase.Translator = translator.AnimalTranslator{}

	model, err := animalUseCase.Fetch()
	if err != nil {
		log.Fatal("Error entity returns nil")
	}

	if model.Name != "Cat" {
		log.Fatal("Error Name value")
	}
}
