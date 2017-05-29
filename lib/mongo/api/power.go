package api

import (

	"gopkg.in/mgo.v2/bson"
	"Go_Server_tripleS/lib/mongo/db"
	"errors"
	"time"
	model "Go_Server_tripleS/lib/mongo/model"
	"math/rand"
)

func UpdatePowerDevice(did string,power int)  error{
	println("LLLLLLL")
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	dev:= model.Device{}

	//Check status
	if err := Db.C("devices").Find(bson.M{"did": did}).One(&dev); err != nil {
		print("Fail")
		return errors.New("No did found")
	}

	year, month, day := time.Now().Date()
	time_now:= time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	colQuerier := bson.M{"did":did,"date": time_now}
	change := bson.M{"$set":
		bson.M{
			"hid": dev.Hid,
			"did": dev.Did,
			"type": dev.Type,
			"dname": dev.Dname,
			"power": power,
			"date": time_now,
			"time" : time_now.Unix(),
		},
	}

	if _,err := Db.C("power").Upsert(colQuerier,change); err != nil {
		print("Fail")
		return errors.New("No did found")
	}

	println("Fukcing wow")
	return nil

}


func InitPowerDevice()  error{

	println("LLLLLLL")
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	year, month, day := time.Now().Date()
	time_now:= time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	if count,err := Db.C("power").Find(bson.M{"date": time_now}).Count(); err != nil {
		println("Fail")
		return errors.New("Eroor")
	}else {
		if count>0{
			println("Existed")
			return errors.New("Eroor")
		}
	}

	var devices []model.DevicePower

	if err := Db.C("devices").Find(nil).All(&devices); err != nil {
		println("Fail")
		return errors.New("Eroor")
	}

	var interfaceSlice []interface{} = make([]interface{}, len(devices))
	for i, dev := range devices {
		dev.Date=time_now
		dev.Time=time_now.Unix()
		dev.Power=0
		interfaceSlice[i] = dev
	}

	if err := Db.C("power").Insert(interfaceSlice...); err != nil {
		println("insert Fail")
		panic(err)
		return errors.New("Eroor")
	}

	println("Fukcing wow")
	return nil

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func InitPowerDeviceTest()  error{

	println("LLLLLLL")
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	var next_day int64=time.Now().Local().Unix()

	for i := 0; i < 1000; i++ {
		next_day-= (24 * 60 * 60 )
		y,m,d:=time.Unix(next_day,0).Date()
		time_now:= time.Date(y, m, d, 0, 0, 0, 0, time.Local)

		if count,err := Db.C("power").Find(bson.M{"date": time_now}).Count(); err != nil {
			println("Fail")
			return errors.New("Eroor")
		}else {
			if count>0{
				println("Existed")
				return errors.New("Eroor")
			}
		}

		var devices []model.DevicePower

		if err := Db.C("devices").Find(nil).All(&devices); err != nil {
			println("Fail")
			return errors.New("Eroor")
		}

		var interfaceSlice []interface{} = make([]interface{}, len(devices))
		for j, dev := range devices {
			dev.Date=time_now
			dev.Time=time_now.Unix()
			dev.Power=0
			interfaceSlice[j] = dev
		}

		if err := Db.C("power").Insert(interfaceSlice...); err != nil {
			println("insert Fail")
			panic(err)
			return errors.New("Eroor")
		}

	}


	println("Fukcing wow")
	return nil

}
