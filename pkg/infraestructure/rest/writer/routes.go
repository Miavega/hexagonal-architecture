package writer

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("/writer", h.PostWriterRoute)
}
