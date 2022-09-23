package controllers

import (
	"net/http"
	"ukiyo/business"

	"github.com/gin-gonic/gin"
)

type PaintingController struct{}

func (p PaintingController) GetPaintingByName(c *gin.Context) {
	name := c.Param("name")
	language := c.Param("language")

	painting := business.GetPaintingByName(name, language)

	c.IndentedJSON(http.StatusOK, painting)
}
