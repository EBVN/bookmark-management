package main

import (
	"github.com/EBVN/bookmark-management/internal/api"
	"github.com/EBVN/bookmark-management/pkg/redis"
)

// @title 		Bookmark Management API
// @version 	1.0.0
// @description API for managing bookmarks
// @BasePath 	/
func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewClient("")
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg, redisClient)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
