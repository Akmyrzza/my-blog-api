package main

import (
	"github.com/Akmyrzza/blog-api/internal/app"
	"github.com/Akmyrzza/blog-api/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%#v", cfg)

	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
