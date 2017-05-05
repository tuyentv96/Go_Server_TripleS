package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	"errors"
)

func UpdatePowerDevice(did string,power int)  error{
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	colQuerier := bson.M{"did": did}
	change := bson.M{"$set": bson.M{"power": power}}

	if err := Db.C("power").Update(colQuerier,change); err != nil {
		print("Fail")
		return errors.New("No did found")
	}

	return nil

}
