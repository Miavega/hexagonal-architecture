package worker

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("/worker", h.PostWorkerRoute)
}
