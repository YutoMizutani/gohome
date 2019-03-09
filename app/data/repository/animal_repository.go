package repository

import "github.com/YutoMizutani/gohome/app/domain/entity"

type AnimalRepository struct {
}

func (repository *AnimalRepository) Fetch() (*entity.Animal, error) {
	animalEntity := &entity.Animal{Name: "Cat"}
	return animalEntity, nil
}
