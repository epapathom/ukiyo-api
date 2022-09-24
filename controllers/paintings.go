package controllers

import (
	"net/http"
	"ukiyo/business"

	"github.com/gin-gonic/gin"
)

type PaintingController struct{}

func (p PaintingController) GetPaintingsByArtist(c *gin.Context) {
	var paintingsQueryParameters PaintingsQueryParameters

	if err := c.ShouldBind(&paintingsQueryParameters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	painting := business.LogicSingleton.GetPaintingsByArtist(paintingsQueryParameters.Artist)

	c.IndentedJSON(http.StatusOK, painting)
}
