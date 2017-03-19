package api

import (
	modelctrl "../../controller/model"
	dbapi "../../mongo/api"
	//	modeldb "../../mongo/model"

	//	"fmt"
	"encoding/json"
	"fmt"
)

func MGetAllDevice(payload []byte)  (modelctrl.Mgethomerespond,bool) {
	m:= modelctrl.Mgethome{}
	bytes:=	[]byte(payload)

	ret:= modelctrl.Mgethomerespond{}

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

	ret.Ltype=data.Ltype
	ret.Ldevice=data.Ldevice
	ret.UID=data.UID
	ret.Uname=data.Uname
	ret.Permission=data.Permission

	fmt.Print("Return value:",ret)

	return ret,false

}