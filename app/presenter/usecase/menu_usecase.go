package usecase

import "github.com/YutoMizutani/gohome/app/domain/model"

type MenuUseCase interface {
	Get() (*model.MenuModels, error)
}
