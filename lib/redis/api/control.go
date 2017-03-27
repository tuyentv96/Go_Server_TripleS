package api

import (
	db "../../redis/db"
	"time"

)

func SaveControlSignalExpire(uid string,cid string,did string,status int)  {
	print("Save to rediss:",uid,cid,"----")
	db.RedCli().Set(did+"S",status,time.Second*5)
	db.RedCli().Set(did+"I",cid,time.Second*5)
	db.RedCli().Set(did+"U",uid,time.Second*5)

	t:= db.RedCli().MGet(did+"S",did+"I").Val()

	for item,index := range t{
		println("\nMGETTTTT:",item,index)
	}

	println("\nMGETHHHHH:",t)

}

func DeviceIsControlling(did string)  bool{
	if val:=db.RedCli().TTL(did+"S").Val().Nanoseconds()/1000000; val>0 {
		return true
	}

	return false

}

func GetControlSignalExpire(did string)  (uid string,cid string,err bool){

	//status= db.RedCli().Get(did+"S").Val()
	uid= db.RedCli().Get(did+"U").Val()
	cid= db.RedCli().Get(did+"I").Val()

	if uid=="" {
		return uid,cid,true
	}

	print(uid,cid)

	return uid,cid,false
}