package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	model "../../mongo/model"

	"fmt"
)

type Param struct {
	Id string
	Pwd string
}


func GetUserByID(data model.User) (model.User,bool){
	id := data.Uid
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.User{}

	if err := Db.C("users").Find(bson.M{"uid": id}).One(&result); err != nil {
		print("Fail")
		return model.User{},true
	}

	print(result.Pwd)
	fmt.Printf("%+v\n",result)
	return result,false


}
