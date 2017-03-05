package api

import (

	modelconn "../../mqtt/model"
	//dbapi "../../mongo/api"

	"fmt"
)

func MLoginHandle(connd modelconn.RqDetail,model modelconn.Mlogin)  {

	fmt.Println("Handle mlogin",connd,model)
	//dbapi.GetUserByID()


}
