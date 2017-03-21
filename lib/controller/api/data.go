package api

import (
	modelctrl "../../controller/model"
	dbapi "../../mongo/api"
	"encoding/json"

)

func MGetAllDevice(payload []byte)  (modelctrl.Mgethomerespond,bool) {
	m:= modelctrl.Mgethome{}
	bytes:=	[]byte(payload)

	ret:= modelctrl.Mgethomerespond{Title:"RMGETDEVICE"}

	err:=	json.Unmarshal(bytes,&m)
	if err!=nil {
		println("parse json fail")
	}

	data,queryerr:= dbapi.MGetAllDevice(m.Uid)

	if queryerr {
		println("Query err")
		ret.Rcode=201
		return ret,true
	}
	ret.Rcode=200
	ret.Lhome=data.Lhome
	ret.Ldevice=data.Ldevice
	ret.UID=data.UID
	ret.Uname=data.Uname
	ret.Permission=data.Permission

	//fmt.Print("Return value:",ret)

	return ret,false

}