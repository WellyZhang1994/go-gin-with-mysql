package server

import "github.com/WellyZhang1994/go-gin-with-mysql/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
