package redis

import (
	"github.com/go-redis/redis"
	"fmt"
	conf "Go_Server_tripleS/conf"
)

var RedisMainCli *redis.Client

type redisDb struct {
	session *redis.Client
}

func InitRedisDb() {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis_server,
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
