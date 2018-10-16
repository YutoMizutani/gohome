package translator

import (
	"github.com/YutoMizutani/gohome/app/data/entity"
	"github.com/YutoMizutani/gohome/app/domain/model"
)

type AnimalTranslator struct {
}

func (translator *AnimalTranslator) Translate(animalEntity *entity.AnimalEntity) (animalModel *model.AnimalModel) {
	animalModel = &model.AnimalModel{Name: animalEntity.Name}
	return animalModel
}
