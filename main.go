package main

import (
	"fmt"
	mqttlib "./lib/mqtt"
	redis "Go_Server_tripleS/lib/redis/db"
	mg "Go_Server_tripleS/lib/mongo/api"
	cron "github.com/robfig/cron"
)

func main() {
	// Create Power Record for new day
	cron := cron.New()
	cron.AddFunc("0 0 0 * * *", func() { mg.InitPowerDevice() })
	cron.Start()

	mg.InitPowerDevice()
	mg.InitPowerDeviceTest()
	fmt.Print("Started server!!!\n")

	redis.InitRedisDb()

	mqttlib.InitMqtt()

	// Loop forever
	select {
	}

}
