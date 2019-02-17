package controller

import (
	"github.com/YutoMizutani/gohome/app/domain/model"
)

type MenuController struct {
}

func (controller *MenuController) Get(c Context) {
	models := model.MenuModels{}

	// Life
	lifeModel := model.MenuModel{
		Title:       "Lives",
		Description: "Daily lives contents",
		Routes:      model.Routes{model.Route{Title: "Weather", Path: "/weather"}},
	}
	models = append(models, lifeModel)

	c.JSON(200, models)
}
