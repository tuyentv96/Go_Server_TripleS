package api

import (

	"gopkg.in/mgo.v2/bson"
	"../../mongo/db"
	"../../mongo/model"

)


func GetUserByID(ctx string) {
	id := ctx.Param("userid")

	Db := db.MgoDb{}
	Db.Init()

	result := model.User{}

	if err := Db.C("people").Find(bson.M{"id": id}).One(&result); err != nil {
		ctx.JSON(iris.StatusOK, model.Err("1"))
		return
	}

	ctx.JSON(iris.StatusOK, &result)

	Db.Close()
}
