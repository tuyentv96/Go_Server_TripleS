package controller

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	model "./model"
//	modeldb "../mongo/model"
//	"encoding/json"
	api "./api"
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

		x:= api.MLoginHandle(info,payload)
		p2,_:= json.Marshal(x)
		x1:=info.Cid
		t:="/RMLOGIN"
		t2:= x1+t
		client.Publish(t2,0,false,p2)

		break
	case "MCONTROL":
		print("Mcontrol handler")
		break

	default:
		print("Topic not found")


	}
}
