package main

import (
	"mginmall/conf"
	"mginmall/routes"
)

func main() {
	conf.InitConfig()
	r := routes.NewRouter()
	_ = r.Run(conf.Config.Service.HttpPort)
}
