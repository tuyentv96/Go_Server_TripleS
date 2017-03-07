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
		value,datquery:= api.MControlHandle(info,payload)
		fmt.Print("valueeee",value)
		if value.Rcode==200 {
			fmt.Printf("%+v", datquery)
			fmt.Println("Continue send to home\n")
			sctrl:= model.Scontrol{Cid:info.Cid,Uid:datquery.Uid,Status:datquery.Status,Did:datquery.Did}
			p2,_:= json.Marshal(sctrl)
			x1:=info.Cid
			t:="/SCONTROL"
			t2:= x1+t
			client.Publish(t2,0,false,p2)


		}
		break

	case "RSCONTROL":
		println("RSCONTROL topic")

		rsp,datquery:= api.MControlRespondHandle(payload)
		fmt.Print(rsp)

		cid:=datquery.Cid
		topic:=cid+"/RMCONTROL"
		client.Publish(topic,0,false,rsp)

	case "MGETDEVICE":




	default:
		print("Topic not found")


	}
}
