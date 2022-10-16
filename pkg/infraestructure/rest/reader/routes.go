package reader

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerEndpoints(rg *gin.RouterGroup) {
	rg.GET("/reader", h.GetRoute)
}
