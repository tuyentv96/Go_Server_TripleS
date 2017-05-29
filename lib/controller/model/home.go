package model


import (
	dbmodel "Go_Server_tripleS/lib/mongo/model"

)

type LHomeDevice struct {
	Hid string `json:"hid"`
}

type LHomeDeviceRsp struct {

	Title string `json:"title"`
	Rcode int `json:"rcode"`
	Ldevice []dbmodel.Device `json:"ldevice"`
}

