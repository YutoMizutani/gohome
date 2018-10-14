package repository

import Entity "gohome/app/data/entity"

type AnimalRepository struct {
}

func (repository *AnimalRepository) Fetch() (entity Entity.AnimalEntity, err error) {
	entity = Entity.AnimalEntity{Name: "Cat"}
	return
}
