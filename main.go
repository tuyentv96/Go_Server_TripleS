package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"time"
	mqttc "./lib/mqtt"

)


func gorou(i int) (int, int) {
	fmt.Printf("Gorou!!!!!%v\n", i)
	time.Sleep(time.Second * 3)
	fmt.Printf("Gorou!!!!!\n")

	return 1, 0
}

func main() {
	go gorou(1)
	mqttc.InitMqtt()


	for true {
	}

}
