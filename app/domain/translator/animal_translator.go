package translator

import (
	"gohome/app/data/entity"
	Model "gohome/app/domain/model"
)

type AnimalTranslator struct {
}

func (translator *AnimalTranslator) Translate(entity *entity.AnimalEntity) (model Model.AnimalModel) {
	model = Model.AnimalModel{Name: entity.Name}
	return model
}
