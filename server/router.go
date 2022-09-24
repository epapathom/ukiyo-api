package server

import (
	"ukiyo/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	healthController := new(controllers.HealthController)
	paintingController := new(controllers.PaintingController)

	router.GET("/", healthController.Status)
	router.GET("/paintings", paintingController.GetPaintingsByArtist)

	return router
}
