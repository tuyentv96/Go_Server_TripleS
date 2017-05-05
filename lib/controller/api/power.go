package api

import (

	modelctrl "../../controller/model"
	dbapi "../../mongo/api"
	"encoding/json"
	"errors"
)


func UpdatePowerDevice(payload []byte)  (modelctrl.UpdatePowerRsp,modelctrl.UpdatePower,error){
	ret:=modelctrl.UpdatePowerRsp{Title:"RUPDATEPOWER",Rcode:200}

	var m modelctrl.UpdatePower
	bytes:=	[]byte(payload)

	if err:=json.Unmarshal(bytes,&m);err!=nil {
		ret.Rcode=201
		return ret,m,errors.New("Wrong format")
	}


	if err:= dbapi.UpdatePowerDevice(m.Did,m.Power);err!=nil {
		ret.Rcode=202
		return ret,m,errors.New("Update power fail")
	}

	return ret,m,nil

}
