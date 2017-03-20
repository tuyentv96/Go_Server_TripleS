package api


import (

	modelctrl "../../controller/model"
	dbapi "../../mongo/api"
//	modeldb "../../mongo/model"

	//	"fmt"
	"encoding/json"
	"fmt"
)

func MControlHandle(ctrlm modelctrl.RqDetail,payload []byte)  (modelctrl.Mcontrolrespond,modelctrl.Mcontrol){
	var m modelctrl.Mcontrol
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelctrl.Mcontrolrespond{Title:"RMCONTROL"}
	ret.Did=m.Did
	ret.Hid=m.Hid

	println("HELOOOO",m.Uid)

	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret,m

	}

	datquery,queryerr:= dbapi.GetDeviceByDid(m.Did)

	fmt.Printf("%+v",datquery)

	if queryerr {
		ret.Rcode=102
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

	fmt.Printf("%+v", datquery)
	ret.Rcode=200
	return ret,m


}

func MControlRespondHandle(payload []byte)  (modelctrl.Mcontrolrespond,modelctrl.Scontrol){

	var m modelctrl.Scontrol
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelctrl.Mcontrolrespond{Title:"RMCONTROL"}
	ret.Did=m.Did
	ret.Hid=m.Hid

	println("HELOOOO",m.Uid)

	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret,m

	}

	updateerr:= dbapi.UpdateStatusDevice(m.Did,m.Status)
	if updateerr {
		ret.Rcode=400
		return ret,m
	}

	ret.Rcode=200
	return ret,m


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

	updateerr:= dbapi.UpdateStatusDevice(m.Did,m.Status)
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