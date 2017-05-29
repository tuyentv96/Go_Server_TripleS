package api

import (

	modelconn "Go_Server_tripleS/lib/controller/model"
	dbapi "Go_Server_tripleS/lib/mongo/api"
	modeldb "Go_Server_tripleS/lib/mongo/model"

//	"fmt"
	"encoding/json"
)

func MLoginHandle(connd modelconn.RqDetail,payload []byte)  modelconn.Mloginrespond{
	var m modeldb.User
	bytes:=	[]byte(payload)

	err:=	json.Unmarshal(bytes,&m)
	ret:=modelconn.Mloginrespond{Title:"RMLOGIN"}

	if err!=nil {
		//fmt.Print("Error json")
		// Wrong format
		ret.Rcode=100
		return ret

	}

	rs,dberr:= dbapi.GetUserByID(m)

	if dberr {
		//User not found
		ret.Rcode=104
		return ret
	}

	//fmt.Printf("%+v", rs)

	if rs.Pwd!=m.Pwd {
		//Password not correct
		ret.Rcode=410
		return ret

	}

	// Login Success
	ret.Rcode=200
	return ret


}
