package main

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Miavega/config"
	"github.com/Miavega/constants"
	"github.com/Miavega/pkg/infraestructure/rest"
	"github.com/Miavega/pkg/infraestructure/rest/reader"
	"github.com/Miavega/pkg/infraestructure/rest/worker"
	"github.com/Miavega/pkg/infraestructure/rest/writer"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(conf.Server.Mode)
	r := gin.Default()

	if err = newHandler(conf, r); err != nil {
		log.Fatal(err)
	}

	r.Run(conf.Server.Host + ":" + conf.Server.Port)
}

func newHandler(conf *config.Config, r *gin.Engine) (err error) {
	rest.GetHealtCheck(r)
	v1 := r.Group(conf.BaseConfig.BasePath)
	switch conf.BaseConfig.Role {
	case constants.WRITE:
		handler := writer.NewHandler(conf)
		handler.API(v1)
	case constants.READ:
		handler := reader.NewHandler(conf)
		handler.API(v1)
	case constants.WORKER:
		handler := worker.NewHandler(conf)
		handler.API(v1)
	default:
		err = errors.New("Invalid role: " + conf.BaseConfig.Role)
	}

	return
}
