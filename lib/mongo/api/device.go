package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	model "../../mongo/model"
)

func GetDeviceByDid(data string)  (model.Device,bool){

	id := data
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Device{}

	if err := Db.C("devices").Find(bson.M{"did": id}).One(&result); err != nil {
		print("Fail")
		return model.Device{},true
	}

	print(result.Did)
	//fmt.Printf("%+v\n",result)
	return result,false
	
}

func GetAllDeviceByHid(hid string)  ([]model.Device,bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := []model.Device{}

	if err := Db.C("devices").Find(bson.M{"hid": hid}).All(&result); err != nil {
		print("Fail")
		return result,true
	}

	//fmt.Printf("%+v\n",result)
	return result,false

}


func GetAllHomeByUid(uid string)  (model.Userpsmdevice,bool){
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

func UpdateStatusDevice(did string,status int)  bool{
	id := did
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	colQuerier := bson.M{"did": id}
	change := bson.M{"$set": bson.M{"status": status}}

	if err := Db.C("devices").Update(colQuerier,change); err != nil {
		print("Fail")
		return true
	}

	return false
}