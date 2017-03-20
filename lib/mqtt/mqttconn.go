package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	handler "../../lib/controller"

	"fmt"
	"os"
	utils "./api"
)




var mqttReceive MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())

	info,errc:= utils.CutTopic(msg.Topic())

	if  errc{
		fmt.Print("Topic not detect")
		return
	}

	fmt.Print("Topic receive: ",info)

	go handler.HandleRequest(client,info,msg.Payload())

}

func InitMqtt()  {
	opts := MQTT.NewClientOptions().AddBroker(utils.Mqttbroker)
	opts.SetClientID(utils.RandStringRunes(14))
	opts.SetDefaultPublishHandler(mqttReceive)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Print("connnect ok")
}


