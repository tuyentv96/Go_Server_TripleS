package main

import (
	"fmt"
	mqttlib "./lib/mqtt"

)


func main() {
	fmt.Print("Started server!!!\n")
	go mqttlib.InitMqtt()

	// Loop forever
	select {
	}

}
