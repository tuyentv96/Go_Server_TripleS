package main

import (
	"fmt"
	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"


	"time"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

type Devices struct {
	Id string
	//ID	bson.ObjectId
	Dname      string
	Address     string
	Status    int
	Power	int
}

func gorou(i int)  (int,int){
	fmt.Printf("Gorou!!!!!%v\n",i)
	time.Sleep(time.Second*3)
	fmt.Printf("Gorou!!!!!\n")
	return  1,0
}

func main() {
	go gorou(1)

	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://trantuyen.name.vn:1883")
	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)
	opts.SetAutoReconnect(true)
	opts.SetWill("tuyen","1234",0,true)

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

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)
	c1:=	session.DB("quanlynha").C("devices")

	// Index
	index := mgo.Index{
		Key:        []string{"ID1"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c1.EnsureIndex(index)
	if err != nil {
		panic(err)
	}



	// Query One
	result:=	Devices{}
	quer:=	bson.M{"_id": bson.ObjectIdHex("587f0c6206722224684c2c08")}
	err = c1.Find(quer).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone ", result,"-\n")



	for true{
	}


	//c.Disconnect(250)
}