package repository

import "gohome/app/data/entity"

type AnimalRepository interface {
	Fetch() (entity.AnimalEntity, error)
}
