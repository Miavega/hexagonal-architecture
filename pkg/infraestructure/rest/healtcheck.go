package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealtCheck(r *gin.Engine) {
	r.GET("/ping", getPing)
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
