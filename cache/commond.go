package cache

import (
	"strconv"

	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	//RedisPw     string
	RedisDbName string
)

func InitRedis(redisAddr, redisDBName string) {
	RedisAddr = redisAddr
	//RedisPw = redisPassword
	RedisDbName = redisDBName
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//Password: RedisPw,
		DB: int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client

}
