package usecase

import "github.com/YutoMizutani/gohome/app/domain/model"

type AnimalUseCase interface {
	Fetch() (*model.AnimalModel, error)
}
