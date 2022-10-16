package reader

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "reader",
	})
}
