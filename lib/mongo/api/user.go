package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	model "../../mongo/model"
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
	//fmt.Printf("%+v\n",result)
	return result,false


}

func CheckPermissonControlDevice(uid string,did string)  bool{

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Userpsmdevice{}

	if err := Db.C("users").Find(bson.M{"uid": uid, "permission": bson.M{"$elemMatch": bson.M{"did": did}}}).One(&result); err != nil {
		print("Fail Check perrrrmsssionnn")
		return false
	}

	//fmt.Printf("Check perrrrmsssionnn ok%+v\n",result)
	return true
}

func GetUserPermissionByID(uid string) (model.Userpsmdevice,bool){

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Userpsmdevice{}

	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&result); err != nil {
		print("Fail")
		return result,true
	}

	//fmt.Printf("%+v\n",result)
	return result,false


}

