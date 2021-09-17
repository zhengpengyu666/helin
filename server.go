package main

import (
	"helin/config"
	"helin/routers"
)

func main() {
	config.Init()
	r := routers.InitRouters()
	if err := r.Start(config.GetConfig().GetString("server.port")); err != nil {
		panic(err)
	}
}
