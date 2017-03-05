package mqtt

import (
	"fmt"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	handler "../../lib/controller"
	model "../controller/model"
	"strings"

)

func demo() {
	fmt.Println("HI")
}

const	(
	mqttbroker="tcp://trantuyen.name.vn:1883"


)

type Person struct {
	Id string `json:"id"`
	Pwd  string `json:"pwd"`
}

type MqttConn struct {
	Conn	*MQTT.Client

}
//define a function for the default message handler


type Devices struct {
	Id string
	//ID	bson.ObjectId
	Dname      string
	Address     string
	Status    int
	Power	int
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


	/*
	bytes:=	[]byte(msg.Payload())

	var p Person


	err:=	json.Unmarshal(bytes,&p)

	if err!=nil {
		fmt.Print("Error json")
		return

	}

	//handler.HandleRequest(client,msg.Topic(),msg.Payload())
	fmt.Print("ID:",p)

	var p1 Person
	json.Unmarshal([]byte(msg.Payload()),&p1)

	fmt.Printf("%+v", p1)
	p2,err:=	json.Marshal(p1)
	fmt.Print(string(p2))
	*/
}

func (w *MqttConn) initMQTT()  MQTT.Client{

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


	return c

}


type MqttC struct {
	Conn MQTT.Client
}

func (w *MqttC) InitMqtt() MQTT.Client {

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

	w.Conn=c

	return c
}

