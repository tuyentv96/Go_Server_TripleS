package api

import (

	"gopkg.in/mgo.v2/bson"
	"Go_Server_tripleS/lib/mongo/db"
	model "Go_Server_tripleS/lib/mongo/model"
	"time"
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

func CheckDeviceStatusByDid(did string,status int)  (bool){

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Device{}

	if err := Db.C("devices").Find(bson.M{"did": did}).One(&result); err != nil {
		// Device id not found
		return true
	}

	if status==result.Status {
		return true
	}

	//Status not match
	return false

}

func GetHomeIDByDid(did string)  (string){

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.Device{}

	if err := Db.C("devices").Find(bson.M{"did": did}).One(&result); err != nil {
		// Device id not found
		return ""
	}

	//Status not match
	return result.Hid

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

func UpdateStatusDevice(did string,status int,uid string)  bool{
	id := did
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	dev:= model.Device{}

	//Check status
	if err := Db.C("devices").Find(bson.M{"did": did}).One(&dev); err != nil {
		print("Fail")
		return true
	}

	if dev.Status==status {
		return true
	}

	//Update status
	colQuerier := bson.M{"did": id}
	change := bson.M{"$set": bson.M{"status": status}}

	if err := Db.C("devices").Update(colQuerier,change); err != nil {
		print("Fail")
		return true
	}
	dev.Status=status
	SaveHistoryDevice(dev,uid)
	return false
}

func SaveHistoryDevice(dev model.Device,uid string)  {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("history").Insert(model.HistoryDevice{Hid:dev.Hid,Did:dev.Did,Status:dev.Status,Dname:dev.Dname,Type:dev.Type,Uid:uid,Time:time.Now().Unix()}); err != nil {
		print("Fail")

	}


}