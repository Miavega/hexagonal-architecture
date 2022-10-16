package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	API(rg *gin.RouterGroup)
}
