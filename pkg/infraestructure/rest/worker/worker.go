package worker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostWorkerRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "worker",
	})
}
