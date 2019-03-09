package controller

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
}

func (controller *MenuController) Get(c *gin.Context) {
	entities := entity.Menus{}

	// Life
	life := entity.Menu{
		Title:       "Lives",
		Description: "Daily lives contents",
		Routes:      entity.Routes{entity.Route{Title: "Weather", Path: "/weather"}},
	}
	entities = append(entities, life)

	c.JSON(200, entities)
}
