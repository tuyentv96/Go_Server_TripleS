package main

import (
	"fmt"
	mqttlib "./lib/mqtt"
	redis "./lib/redis/db"
	//redisapi "./lib/redis/api"

)


func main() {
	fmt.Print("Started server!!!\n")

	redis.InitRedisDb()
	//redisapi.SaveControlSignalExpire("d1",1)


	mqttlib.InitMqtt()

	// Loop forever
	select {
	}

}
