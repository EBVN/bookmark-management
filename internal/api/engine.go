package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EBVN/golang-course/internal/handler"
	"github.com/EBVN/golang-course/internal/service"
)

type Engine interface {
	Run() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

func NewEngine(cfg *Config) Engine {
	api := &engine{
		app: gin.New(),
		cfg: cfg,
	}

	api.initRoutes()
	return api
}

func (e *engine) Run() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

func (e *engine) initRoutes() {
	// create handler
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	e.app.GET("/genpass", genPassHandler.GeneratePassword)
}
