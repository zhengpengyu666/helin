package main

import (
	"fmt"
	"helin/config"
	"helin/routers"
)

func main() {
	config.Init()
	r := routers.InitRouters()
	if err := r.Start(config.GetConfig().GetString("server.port")); err != nil {
		fmt.Println("启动异常")
	}
}
