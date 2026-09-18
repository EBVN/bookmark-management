package main

import (
	"github.com/EBVN/bookmark-management/internal/api"
	"github.com/EBVN/bookmark-management/pkg/logger"
	"github.com/EBVN/bookmark-management/pkg/redis"
)

// @title 		Bookmark Management API
// @version 	1.0.0
// @description API for managing bookmarks
// @BasePath 	/
func main() {
	// get app config
	appConfig, err := api.NewConfig("")
	if err != nil {
		panic(err)
	}

	logger.SetLogLevel(appConfig.LogLevel)

	redisClient, err := redis.NewClient("")
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	app := api.NewEngine(appConfig, redisClient)

	err = app.Start()
	if err != nil {
		panic(err)
	}
}
