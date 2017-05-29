package controller

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	model "Go_Server_tripleS/lib/controller/model"
	api "Go_Server_tripleS/lib/controller/api"
	"encoding/json"
	//mqttlib "../mqtt/api"

//	"github.com/go-redis/redis"
)



func  HandleRequest(client MQTT.Client,info model.RqDetail,payload []byte)   {

	switch info.Topic {

	case "MLOGIN":
		println("Login handler")
		//client.Publish

		rsp:= api.MLoginHandle(info,payload)
		payl,_:= json.Marshal(rsp)
		rtopic:=info.Cid+"/RMLOGIN"
		client.Publish(rtopic,0,false,payl)

		break
	case "MCONTROL":
		print("Mcontrol handler")
		value,datquery:= api.MControlHandle(info,payload)
		if value.Rcode==200 {


			//fmt.Printf("%+v", datquery)
			//fmt.Println("Continue send to home\n")
			sctrl:= model.Scontrol{Status:datquery.Status,Did:datquery.Did,Hid:datquery.Hid}
			payl,_:= json.Marshal(sctrl)

			client.Publish(datquery.Hid+"/SCONTROL",0,false,payl)


		} else {
			payl,_:= json.Marshal(value)
			client.Publish(info.Cid+"/RMCONTROL",0,false,payl)

		}
		break

	case "MCONTROLS":
		print("McontrolS handler")
		data:= api.MControlsHandle(payload)
		payl,_:= json.Marshal(data)
		client.Publish(data.Hid+"/SCONTROLS",0,false,payl)
		break
/*
	case "RSCONTROL":
		println("RSCONTROL topic")

		rsp,datquery,uid,cid:= api.MControlRespondHandle(payload)
		fmt.Print("uiddddd",uid)
//		redis.GetControlSignalExpire(datquery.Did)

		payl,_:= json.Marshal(rsp)
		topic:=cid+"/RMCONTROL"
		client.Publish(topic,0,false,payl)

		msyncdata:= api.MSync(datquery.Hid,datquery.Did,datquery.Status)

		paylsync,_:=json.Marshal(msyncdata)
		topicsync:=datquery.Hid+"/MSYNC"
		client.Publish(topicsync,0,false,paylsync)
*/

	case "CONTROL":
		ctr_type:= api.Check_Type_Control(payload)
		println("Control type:",ctr_type)
		if ctr_type==1 {
			rsp, datquery, uid, cid := api.MControlRespondHandle(payload)
			fmt.Print("mcontrol", uid)
			//		redis.GetControlSignalExpire(datquery.Did)

			payl, _ := json.Marshal(rsp)
			topic := cid + "/RMCONTROL"
			client.Publish(topic, 0, false, payl)

			msyncdata:= api.MSync(datquery.Hid,datquery.Did,datquery.Status)

			paylsync,_:=json.Marshal(msyncdata)
			topicsync:=datquery.Hid+"/MSYNC"

			//Sync to mobile
			if rsp.Rcode==200 {
				client.Publish(topicsync,0,false,paylsync)
			}
		}else {

			rsp, datquery := api.ControlDevice(payload)
			fmt.Print("ctonrol",rsp,datquery)


			payl,_ := json.Marshal(rsp)
			cid:= info.Cid
			topic := cid+"/RCONTROL"
			client.Publish(topic, 0, false, payl)

			msyncdata:= api.MSync(datquery.Hid,datquery.Did,datquery.Status)

			paylsync,_:=json.Marshal(msyncdata)
			topicsync:=datquery.Hid+"/MSYNC"

			//Sync to mobile
			if rsp.Rcode==200 {
				client.Publish(topicsync,0,false,paylsync)
			}
		}



	case "SCONTROL":
		client.Publish(info.Cid+"/CONTROL",0,false,payload)


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

	case "MGETDEVICE":
		println("mgetdevice !!")
		rsp,_:= api.MGetAllDevice(payload)

		payl,_ := json.Marshal(rsp)

		client.Publish(info.Cid+"/RMGETDEVICE",0,false,payl)

	case "GETDEVICE":
		println("getdevice !!!")
		rsp:= api.HomeGetAllDevice(payload)

		payl,_ := json.Marshal(rsp)

		client.Publish(info.Cid+"/RGETDEVICE",0,false,payl)

	case "UPDATEPOWER":
		println("update power !!!")
		rsp,data,err:= api.UpdatePowerDevice(payload)

		pay,_ := json.Marshal(rsp)
		client.Publish(info.Cid+"/RUPDATEPOWER",0,false,pay)

		if err!=nil {
			return
		}

		payl,_ := json.Marshal(model.MSyncPower{Title:"MSYNCPOWER",Hid:data.Hid,Did:data.Did,Power:data.Power})

		client.Publish(info.Cid+"/MSYNCPOWER",0,false,payl)



	default:
		print("Topic not found")
	}


	return
}
