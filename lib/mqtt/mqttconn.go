package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	handler "../../lib/controller"
	model "../../lib/controller/model"
	"fmt"
	utils "./api"
	"os"
)

var topics  = map[string]byte{

	"+/MLOGIN": 0,
	"+/MCONTROL":   0,
	"+/RSCONTROL": 0,
	"+/CONTROL": 0,
	"+/SCONTROL": 0,
	"+/HGOOFF":   0,
	"+/HGOON": 0,
	"+/MGETDEVICE": 0,

}




var mqttReceive MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())
	info:= model.RqDetail{}

	if err:= utils.CutTopic(msg.Topic(),&info); err{
		fmt.Print("Topic not detect")
		return
	}

	fmt.Print("Topic receive: ",info)

	go handler.HandleRequest(client,info,msg.Payload())

}

func InitMqtt()  {
	opts := MQTT.NewClientOptions().AddBroker(utils.Mqttbroker)
	opts.SetClientID(utils.RandStringRunes(10))
	opts.SetDefaultPublishHandler(mqttReceive)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	/*
	if token := c.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	*/

	if token :=c.SubscribeMultiple(topics,nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}


	fmt.Print("connnect ok")
}


