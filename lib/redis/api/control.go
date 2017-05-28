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

func GetControlSignalExpireNoDel(did string)  (uid string,cid string,err bool){

	//status= db.RedCli().Get(did+"S").Val()
	uid= db.RedCli().Get(did+"U").Val()
	cid= db.RedCli().Get(did+"I").Val()

	if cid=="" {
		return uid,cid,true
	}


	print(uid,cid)

	return uid,cid,false
}

func GetControlSignalExpire(did string)  (uid string,cid string,err bool){

	//status= db.RedCli().Get(did+"S").Val()
	uid= db.RedCli().Get(did+"U").Val()
	cid= db.RedCli().Get(did+"I").Val()

	if cid=="" {
		return uid,cid,true
	}

	db.RedCli().Del(did+"S")
	db.RedCli().Del(did+"U")
	db.RedCli().Del(did+"I")

	print(uid,cid)

	return uid,cid,false
}

func TimerSaveControlSignalExpire(uid string,did string,status int)  {
	print("Save to rediss:",uid,"----")
	db.RedCli().Set(did+"S",status,time.Second*5)
	db.RedCli().Set(did+"U",uid,time.Second*5)
	/*
	t:= db.RedCli().MGet(did+"S",did+"I").Val()

	for item,index := range t{
		println("\nMGETTTTT:",item,index)
	}

	println("\nMGETHHHHH:",t)
	*/

}

func TimerDeviceIsControlling(did string)  bool{
	if val:=db.RedCli().TTL(did+"S").Val().Nanoseconds()/1000000; val>0 {
		return true
	}

	return false

}

func TimerGetControlSignalExpire(did string)  (uid string,err bool){

	//status= db.RedCli().Get(did+"S").Val()
	uid= db.RedCli().Get(did+"U").Val()

	if uid=="" {
		return uid,true
	}

	db.RedCli().Del(did+"S")
	db.RedCli().Del(did+"U")

	print(uid)

	return uid,false
}