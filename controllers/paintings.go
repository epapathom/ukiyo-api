package controllers

import (
	"net/http"
	"ukiyo/business"

	"github.com/gin-gonic/gin"
)

type PaintingController struct{}

// GetPaintingsByArtist godoc
// @Summary      Get paintings
// @Description  Get paintings
// @Tags         paintings
// @Accept       json
// @Produce      json
// @Param        artist  query  string  false  "Search by artist"
// @Success      200  {array}  models.Painting
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Router       /paintings [get]
func (p PaintingController) GetPaintingsByArtist(c *gin.Context) {
	var paintingsQueryParameters PaintingsQueryParameters

	if err := c.ShouldBind(&paintingsQueryParameters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	painting := business.LogicSingleton.GetPaintingsByArtist(paintingsQueryParameters.Artist)

	c.IndentedJSON(http.StatusOK, painting)
}
