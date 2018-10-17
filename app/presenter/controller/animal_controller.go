package controller

import "github.com/YutoMizutani/gohome/app/presenter/usecase"

type AnimalController struct {
	UseCase usecase.AnimalUseCase
}

func (controller *AnimalController) Fetch(c Context) {
	animals, err := controller.UseCase.Fetch()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, animals)
}
