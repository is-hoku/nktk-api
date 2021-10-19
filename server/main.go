package main

import (
	"github.com/is-hoku/nktk-api/server/config"
	"github.com/is-hoku/nktk-api/server/router"
)

func main() {
	config.Init("config.yml")
	r := router.NewRouter()
	//	r.Run(config.Config.Port)
	r.Logger.Fatal(r.Start(config.Config.Port))
}
