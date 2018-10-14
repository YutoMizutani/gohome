package infrastructure

import (
	Builder "gohome/app/application/builder"

	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	builder := Builder.AnimalBuilder{}
	animalController := builder.Build()

	router.GET("/animal", func(c *gin.Context) { animalController.Fetch(c) })

	Router = router
}
