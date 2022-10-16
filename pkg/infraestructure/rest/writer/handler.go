package writer

import (
	"github.com/gin-gonic/gin"

	"github.com/Miavega/config"
	"github.com/Miavega/pkg/infraestructure/rest"
)

type Handler struct {
	Config *config.Config
}

func NewHandler(conf *config.Config) rest.Handler {
	return &Handler{
		Config: conf,
	}
}

func (h *Handler) API(rg *gin.RouterGroup) {
	h.registerEndpoints(rg)
}
