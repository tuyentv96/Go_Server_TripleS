package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	model "../controller/model"
	"strings"
	handler "../../lib/controller"
	"math/rand"

	"fmt"
	"os"
)


const	(
	Mqttbroker="tcp://trantuyen.name.vn:1883"
)


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MqttPublishSingle(topic string,qos byte,retain bool,payload []byte)  {
	opts := MQTT.NewClientOptions().AddBroker(Mqttbroker)
	opts.SetClientID(RandStringRunes(12))
	opts.SetCleanSession(true)


	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token:= c.Publish(topic,qos,retain,payload)
	token.Wait()
	c.Disconnect(0)
}

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

func InitMqtt()  {
	opts := MQTT.NewClientOptions().AddBroker(Mqttbroker)
	opts.SetClientID(RandStringRunes(12))
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
}


