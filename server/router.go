package server

import (
	"ukiyo/controllers"
	"ukiyo/utils"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	healthController := new(controllers.HealthController)
	paintingController := new(controllers.PaintingController)

	router.GET(utils.AddStageToPath("/"), healthController.Status)
	router.GET(utils.AddStageToPath("/paintings/:artist"), paintingController.GetPaintingsByArtist)

	return router
}
