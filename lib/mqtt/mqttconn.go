package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	handler "../../lib/controller"
	model "../../lib/controller/model"
	"fmt"
	utils "./api"
	"os"
	"time"
	conf "Go_Server_tripleS/conf"
)

var topics  = map[string]byte{

	"+/MLOGIN": 0,
	"+/MCONTROL":   0,
	"+/MCONTROLS":   0,
	"+/RSCONTROL": 0,
	"+/CONTROL": 0,
	"+/SCONTROL": 0,
	"+/HGOOFF":   0,
	"+/HGOON": 0,
	"+/MGETDEVICE": 0,
	"+/GETDEVICE": 0,
	"+/UPDATEPOWER": 0,

}

var MainMqttC MainMqttClient

type MainMqttClient struct {
	client MQTT.Client
	opt *MQTT.ClientOptions
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

func mqttOnConnect(client MQTT.Client) {
	print("connected to broker")
}

func mqttLostConnect(client MQTT.Client,err error) {
	print("disconnect to broker")


}




func InitMqtt()  {
	cid:= string(string(time.Now().Unix())+utils.RandStringRunes(12))
	print("cid:",cid)
	opts := MQTT.NewClientOptions().AddBroker(conf.Mqtt_broker)
	opts.SetClientID(cid)
	opts.SetDefaultPublishHandler(mqttReceive)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)
	opts.SetOnConnectHandler(mqttOnConnect)
	opts.SetConnectionLostHandler(mqttLostConnect)


	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	if token :=c.SubscribeMultiple(topics,nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	MainMqttC.client=c
	fmt.Print("connnect ok")


}



