package usecase

import "github.com/YutoMizutani/gohome/app/domain/entity"

type AnimalUseCase interface {
	Fetch() (*entity.Animal, error)
}
