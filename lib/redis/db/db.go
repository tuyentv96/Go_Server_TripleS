package redis

import (
	"github.com/go-redis/redis"
	"fmt"
)

var RedisMainCli *redis.Client

type redisDb struct {
	session *redis.Client
}

func InitRedisDb() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:32768",
		Password: "", // no password set
		DB:       0,  // use default DB

	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	RedisMainCli=client
}

func TestRedis()  {
	pong,err :=RedisMainCli.Ping().Result()
	fmt.Println(pong, err)
}

func RedCli()  *redis.Client{
	return RedisMainCli
}
