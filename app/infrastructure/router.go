package infrastructure

import (
	"github.com/YutoMizutani/gohome/app/application/builder"

	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	animalBuilder := builder.AnimalBuilder{}
	animalController := animalBuilder.Build()

	router.GET("/animal", func(c *gin.Context) { animalController.Fetch(c) })

	Router = router
}
