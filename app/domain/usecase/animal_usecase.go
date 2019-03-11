package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/repository"
)

type AnimalUseCase struct {
	Repository repository.AnimalRepository
}

func (usecase *AnimalUseCase) Fetch() (*entity.Animal, error) {
	entity, err := usecase.Repository.Fetch()
	if err != nil {
		return nil, err
	}

	return entity, nil
}
