package api


import (

	modelctrl "../../controller/model"
	dbapi "../../mongo/api"
	"encoding/json"
	redis "../../redis/api"
	"fmt"
)

func MControlHandle(ctrlm modelctrl.RqDetail,payload []byte)  (modelctrl.Mcontrolrespond,modelctrl.Mcontrol){
	var m modelctrl.Mcontrol
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelctrl.Mcontrolrespond{Title:"RMCONTROL"}
	ret.Did=m.Did
	ret.Hid=m.Hid
	ret.Status=m.Status

	//println("HELOOOO",m.Uid)

	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret,m

	}
	// Check device is ready to control

	if redis.DeviceIsControlling(m.Did){
		ret.Rcode=155
		return ret,m
	}

	//Check permission for user to control device
	if dbapi.CheckPermissonControlDevice(m.Uid,m.Did)==false {
		ret.Rcode=111
		return ret,m
	}


	datquery,queryerr:= dbapi.GetDeviceByDid(m.Did)

	//fmt.Printf("%+v",datquery)

	if queryerr {
		ret.Rcode=102
		return ret,m
	}

	if !(m.Status>=0 && m.Status<=1) {
		ret.Rcode=100
		return ret,m

	}

	if datquery.Status==m.Status {
		ret.Rcode=311
		return ret,m

	}

	if datquery.Hid!=m.Hid {
		ret.Rcode=112
		return ret,m

	}

	redis.SaveControlSignalExpire(m.Uid,ctrlm.Cid,datquery.Did,datquery.Status)

	//fmt.Printf("%+v", datquery)
	ret.Rcode=200
	return ret,m


}

func Check_Type_Control(payload []byte)  int{
	var m modelctrl.Control
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)


	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		return 0

	}

	uid,cid,rderr:= redis.GetControlSignalExpire(m.Did)

	print("uidneee",uid,cid,rderr)

	if rderr {
		return 2
	}

	return 1
}

func MControlRespondHandle(payload []byte)  (modelctrl.Mcontrolrespond,modelctrl.Scontrol,string,string){

	var m modelctrl.Scontrol
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelctrl.Mcontrolrespond{Title:"RMCONTROL"}
	ret.Did=m.Did
	ret.Hid=m.Hid
	ret.Status=m.Status



	if err!=nil {
		fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret,m,"",""

	}
	println("HELOOOO")
	uid,cid,rderr:= redis.GetControlSignalExpire(m.Did)

	print("uidneee",uid,cid)

	if rderr {
		ret.Rcode=400
		return ret,m,"",""
	}

	if !(m.Status>=0 && m.Status<=1) {
		ret.Rcode=100
		return ret,m,"",""

	}

	updateerr:= dbapi.UpdateStatusDevice(m.Did,m.Status,uid)
	if updateerr {
		ret.Rcode=400
		return ret,m,"",""
	}

	redis.GetControlSignalExpire(m.Did)

	ret.Rcode=200
	return ret,m,uid,cid


}

func ControlDevice(payload []byte)  (modelctrl.Controlrsp,modelctrl.Control){

	var m modelctrl.Control
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelctrl.Controlrsp{Title:"RCONTROL"}


	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret,m

	}


	if !(m.Status>=0 && m.Status<=1) {
		ret.Rcode=100
		return ret,m

	}

	updateerr:= dbapi.UpdateStatusDevice(m.Did,m.Status,"root")
	if updateerr {
		ret.Rcode=400
		return ret,m
	}

	ret.Rcode=200
	return ret,m


}


func MSync(hid string,did string,status int)  modelctrl.Msync{
	ret:=modelctrl.Msync{Title:"MSYNC"}
	ret.Hid=hid
	ret.Did=did
	ret.Status=status

	return ret
}