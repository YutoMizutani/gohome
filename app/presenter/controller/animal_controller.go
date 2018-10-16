package controller

import "github.com/YutoMizutani/gohome/app/presenter/usecase"

type AnimalController struct {
	Usecase usecase.AnimalUsecase
}

func (controller *AnimalController) Fetch(c Context) {
	animals, err := controller.Usecase.Fetch()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, animals)
}
