package api

import (
	"fmt"
	"net/http"

	"github.com/EBVN/bookmark-management/internal/handler"
	"github.com/EBVN/bookmark-management/internal/repository"
	"github.com/EBVN/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/EBVN/bookmark-management/docs"
)

// Engine is the interface of engine struct
type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app         *gin.Engine
	cfg         *Config
	redisClient *redis.Client
}

// NewEngine create API Engine
func NewEngine(cfg *Config, redisClient *redis.Client) Engine {
	e := &engine{app: gin.Default(), cfg: cfg, redisClient: redisClient}
	e.initRoutes()
	return e
}

// Start starts the HTTP server on the configured port
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%d", e.cfg.AppPort))
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

func (e *engine) initRoutes() {
	// khoi tao handler
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	urlStorageRepo := repository.NewURLStorage(e.redisClient)
	shortenURLSvc := service.NewShortenUrl(urlStorageRepo, genPassSvc)
	shortenHandler := handler.NewShortenLink(shortenURLSvc)

	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.app.GET("/genpass", genPassHandler.GeneratePassword)
	e.app.POST("/v1/links/shorten", shortenHandler.CreateShortenLink)
	e.app.GET("/v1/links/redirect/:code", shortenHandler.Redirect)

}
