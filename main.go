package main

import (
	"fmt"
	mqttlib "./lib/mqtt"

)


func main() {
	fmt.Print("Started server!!!\n")
	mqttlib.InitMqtt()

	// Loop forever
	select {
	}

}
