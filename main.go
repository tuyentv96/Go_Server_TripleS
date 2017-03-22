package main

import (
	"fmt"
	mqttlib "./lib/mqtt"
	//redis "./lib/redis/db"
	//redisapi "./lib/redis/api"

)


func main() {
	fmt.Print("Started server!!!\n")
	/*
	redis.InitRedisDb()
	for i:=0;i<4000;i++ {
		redisapi.SaveControlSignalExpire()
		//redis.TestRedis()
	}
	*/

	mqttlib.InitMqtt()

	// Loop forever
	select {
	}

}
