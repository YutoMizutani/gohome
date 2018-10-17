package repository

import "github.com/YutoMizutani/gohome/app/data/entity"

type AnimalRepository struct {
}

func (repository *AnimalRepository) Fetch() (*entity.AnimalEntity, error) {
	animalEntity := &entity.AnimalEntity{Name: "Cat"}
	return animalEntity, nil
}
