package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"time"
	mqttc "./lib/mqtt"

)

type Person struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`

}

//define a function for the default message handler

type Devices struct {
	Id string
	//ID	bson.ObjectId
	Dname   string
	Address string
	Status  int
	Power   int
}

func gorou(i int) (int, int) {
	fmt.Printf("Gorou!!!!!%v\n", i)
	time.Sleep(time.Second * 3)
	fmt.Printf("Gorou!!!!!\n")

	return 1, 0
}

func main() {
	go gorou(1)
	Conn:=	mqttc.MqttC{}
	Conn.InitMqtt()


	/*
		//create a ClientOptions struct setting the broker address, clientid, turn
		//off trace output and set the default message handler


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
	*/

	for true {
	}

	//c.Disconnect(250)
}
