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

		rsp:= api.MLoginHandle(info,payload)
		payl,_:= json.Marshal(rsp)
		rtopic:=info.Cid+"/RMLOGIN"
		client.Publish(rtopic,0,false,payl)

		break
	case "MCONTROL":
		print("Mcontrol handler")
		value,datquery:= api.MControlHandle(info,payload)
		if value.Rcode==200 {
			fmt.Printf("%+v", datquery)
			fmt.Println("Continue send to home\n")
			sctrl:= model.Scontrol{Cid:info.Cid,Uid:datquery.Uid,Status:datquery.Status,Did:datquery.Did,Hid:datquery.Hid}
			payl,_:= json.Marshal(sctrl)
			client.Publish(datquery.Hid+"/SCONTROL",0,false,payl)


		} else {
			payl,_:= json.Marshal(value)
			client.Publish(info.Cid+"/RMCONTROL",0,false,payl)

		}
		break

	case "RSCONTROL":
		println("RSCONTROL topic")

		rsp,datquery:= api.MControlRespondHandle(payload)
		fmt.Print(rsp)

		payl,_:= json.Marshal(rsp)
		cid:=datquery.Cid
		topic:=cid+"/RMCONTROL"
		client.Publish(topic,0,false,payl)

		paylsync,_:=json.Marshal(datquery)
		topicsync:=datquery.Hid+"/MSYNC"
		client.Publish(topicsync,0,false,paylsync)


	case "CONTROL":
		rsp,datquery:= api.ControlDevice(payload)
		fmt.Print(rsp)

		payl,_:= json.Marshal(rsp)
		cid:=info.Cid
		topic:=cid+"/RCONTROL"
		client.Publish(topic,0,false,payl)

		//Sync to mobile
		if rsp.Rcode==200 {
			client.Publish(datquery.Hid+"/MSYNC",0,false,payload)
		}
	/*
	case "SCONTROL":
		client.Publish(info.Cid+"/RSCONTROL",0,false,payload)
		*/

	case "HGOOFF":
		println("home go off detected")
		rsp:= api.HomeStatusRev(payload)
		rsp.Status=0
		payl,_ := json.Marshal(rsp)
		client.Publish(rsp.Hid+"/MHSTATUS",0,false,payl)

	case "HGOON":
		println("home go on detected")
		rsp:= api.HomeStatusRev(payload)
		rsp.Status=1
		payl,_ := json.Marshal(rsp)
		client.Publish(rsp.Hid+"/MHSTATUS",0,false,payl)


	default:
		print("Topic not found")


	}
}
