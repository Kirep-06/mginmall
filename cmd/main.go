package main

import (
	"mginmall/cache"
	"mginmall/conf"
	"mginmall/routes"
)

func main() {
	conf.InitConfig()
	cache.InitRedis(
		conf.Config.Redis.RedisAddr,
		conf.Config.Redis.RedisPassword,
		conf.Config.Redis.RedisDbName,
	)
	cache.Redis()
	r := routes.NewRouter()
	_ = r.Run(conf.Config.Service.HttpPort)
}
