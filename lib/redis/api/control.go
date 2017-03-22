package api

import db "../../redis/db"

func SaveControlSignalExpire()  {
	x,_:=db.RedCli().Ping().Result()
	print(x)
}