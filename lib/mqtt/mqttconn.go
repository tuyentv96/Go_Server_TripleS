package mqtt

import (
	"fmt"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	handler "../../lib/controller"
	model "../controller/model"
	"strings"

)


const	(
	mqttbroker="tcp://trantuyen.name.vn:1883"


)


func CutTopic(topic string)  (info model.RqDetail,b bool){
	index:= strings.Index(topic,"/")
	if  index>-1 {
		info.Topic=topic[index+1:]
		info.Cid=topic[:index]
		return info,false
	}
	return info,true

}

var mqttReceive MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())

	info,errc:= CutTopic(msg.Topic())

	if  errc{
		fmt.Print("Topic not detect")
		return
	}

	fmt.Print("Topic receive: ",info)

	go handler.HandleRequest(client,info,msg.Payload())

}


func InitMqtt() MQTT.Client {

	opts := MQTT.NewClientOptions().AddBroker(mqttbroker)
	opts.SetClientID("Servertriples133223")
	opts.SetDefaultPublishHandler(mqttReceive)
	opts.SetAutoReconnect(true)
	//opts.SetWill("tuyen","1234",0,true)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	c.Publish("test",0,false,"nhucccccc")

	return c
}

