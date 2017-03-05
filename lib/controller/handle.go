package controller

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	model "./model"
	"encoding/json"
)

const	(
	mqttbroker="tcp://trantuyen.name.vn:1883"


)


func  HandleRequest(client MQTT.Client,info model.RqDetail,payload []byte)   {

	fmt.Println("FFFF")
	switch info.Topic {

	case "MLOGIN":
		println("Login handler")
		var m model.Mlogin
		bytes:=	[]byte(payload)

		err:=	json.Unmarshal(bytes,&m)

		if err!=nil {
			fmt.Print("Error json")
			return

		}

		fmt.Print("Mogin",m)
		break
	case "MCONTROL":
		print("Mcontrol handler")
		break

	default:
		print("Topic not found")


	}
}
