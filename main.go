package main

import (
	"github.com/EvgenKostenko/stackoverlow_performance/api"
	"github.com/EvgenKostenko/stackoverlow_performance/config"
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/topusers/:word", api.GetTopUsersByWord)
	iris.Get("/toptags/:userid", api.GetTopUsersTags)
	iris.Listen(config.Config.API.Host)
}
