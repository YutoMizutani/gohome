package usecase_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
)

type animalRepositoryMock struct {
}

func (repository *animalRepositoryMock) Fetch() (*entity.Animal, error) {
	animalEntity := &entity.Animal{}
	animalEntity.Name = "Cat"
	return animalEntity, nil
}

func TestAnimalUseCaseFetch(t *testing.T) {
	mock := new(animalRepositoryMock)

	animalUseCase := usecase.AnimalUseCase{}
	animalUseCase.Repository = mock

	entity, err := animalUseCase.Fetch()
	if err != nil {
		log.Fatal("Error entity returns nil")
	}

	if entity.Name != "Cat" {
		log.Fatal("Error Name value")
	}
}
