package repository

import "github.com/YutoMizutani/gohome/app/domain/entity"

type AnimalRepository interface {
	Fetch() (*entity.Animal, error)
}
