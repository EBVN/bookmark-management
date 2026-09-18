package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/EBVN/bookmark-management/docs"
	"github.com/EBVN/bookmark-management/internal/repository"

	"github.com/EBVN/bookmark-management/internal/handler"
	"github.com/EBVN/bookmark-management/internal/service"
)

type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app         *gin.Engine
	cfg         *Config
	redisClient *redis.Client
}

func NewEngine(cfg *Config, redisClient *redis.Client) Engine {
	app := &engine{
		app:         gin.Default(),
		cfg:         cfg,
		redisClient: redisClient,
	}
	app.initRoutes()
	return app
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

func (e *engine) initRoutes() {
	// register handler

	// khoi tao handler
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	urlStorage := repository.NewURLStorage(e.redisClient)
	shortenUrlSvc := service.NewShortenUrl(urlStorage, genPassSvc)
	shortenUrlHandler := handler.NewShortenLink(shortenUrlSvc)

	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.app.GET("/genpass", genPassHandler.GeneratePassword)
	e.app.POST("/v1/links/shorten", shortenUrlHandler.CreateShortenLink)
	e.app.GET("/v1/links/redirect/:code", shortenUrlHandler.Redirect)

}
