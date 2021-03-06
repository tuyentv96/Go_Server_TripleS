package api

import (

	"gopkg.in/mgo.v2/bson"
	"Go_Server_tripleS/lib/mongo/db"
	model "Go_Server_tripleS/lib/mongo/model"
)

func MGetAllDevice(data string)  (model.LDevice,bool){

	id := data
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := model.LDevice{}


	if err := Db.C("users").Find(bson.M{"uid": id}).One(&result); err != nil {
		print("Fail")
		return result,true

	}

	lh:= []bson.M{}

	for i:=0;i<len(result.Lhome);i++ {
		temp:= bson.M{"hid":result.Lhome[i].Hid}
		lh=append(lh,temp)
		
	}

	//fmt.Print("BSSS",lh)
	if err1 := Db.C("devices").Find(bson.M{"$or": lh}).Sort("hid","dname").All(&result.Ldevice); err1 != nil {
		print("Fail")
		return result,true

	}

	//result1,_:=json.Marshal(result)
	//fmt.Print(string(result1))

	return result,false

}

func GetAllDevice(hid string)  (ldevice []model.Device,err bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	if err := Db.C("devices").Find(bson.M{"hid": hid}).Sort("dname").All(&ldevice); err != nil {
		print("Fail")
		return ldevice,true

	}

	return ldevice,false

}