package controller

import "github.com/YutoMizutani/gohome/app/presenter/usecase"

type MenuController struct {
	UseCase usecase.MenuUseCase
}

func (controller *MenuController) Get(c Context) {
	menus, err := controller.UseCase.Get()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, menus)
}
