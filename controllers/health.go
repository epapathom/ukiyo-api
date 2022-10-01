package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status godoc
// @Summary      Health check
// @Description  Health check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string} string
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Router       / [get]
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
