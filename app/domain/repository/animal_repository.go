package repository

import "github.com/YutoMizutani/gohome/app/data/entity"

type AnimalRepository interface {
	Fetch() (entity.AnimalEntity, error)
}
