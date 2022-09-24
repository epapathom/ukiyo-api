package controllers

import (
	"net/http"
	"ukiyo/bindings"
	"ukiyo/business"

	"github.com/gin-gonic/gin"
)

type PaintingController struct{}

func (p PaintingController) GetPaintingsByArtist(c *gin.Context) {
	var paintingsPathParameters bindings.PaintingsPathParameters

	if err := c.ShouldBindUri(&paintingsPathParameters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	logic := new(business.Logic)
	painting := logic.GetPaintingsByArtist(paintingsPathParameters.Artist)

	c.IndentedJSON(http.StatusOK, painting)
}
