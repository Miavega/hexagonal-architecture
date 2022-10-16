package writer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostWriterRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "writer",
	})
}
