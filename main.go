package main

import (
	"github.com/EvgenKostenko/stackoverlow_performance/api"
	"github.com/kataras/iris"
	"github.com/EvgenKostenko/stackoverlow_performance/config"
	"fmt"
)

func main() {
	fmt.Printf("config: %#v", config.Config)
	iris.Get("/topusers/:word", api.GetTopUsersByWord)
	iris.Get("/toptags/:userid", api.GetTopUsersTags)
	iris.Listen(config.Config.API.Host)
}
