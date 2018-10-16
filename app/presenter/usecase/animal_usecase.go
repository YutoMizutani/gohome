package usecase

import "github.com/YutoMizutani/gohome/app/domain/model"

type AnimalUsecase interface {
	Fetch() (*model.AnimalModel, error)
}
