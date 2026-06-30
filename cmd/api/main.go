package main

import "github.com/EBVN/golang-course/internal/api"

func main() {
	// get app config
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}

	apiEngine := api.NewEngine(cfg)
	err = apiEngine.Run()
	if err != nil {
		panic(err)
	}
}
