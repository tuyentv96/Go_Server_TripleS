package model

import (
	dbmodel "../../mongo/model"

)

type Mgethome struct {
	Uid string `json:"uid"`
}

type Mgethomerespond struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
	UID   string `json:"uid"`
	Uname string `json:"uname"`
	Types []struct {
		Hid  string `json:"hid"`
		Type int    `json:"type"`
	} `json:"type"`
	Ldevice []dbmodel.Device `json:"ldevice"`
	Permission []struct {
		Hid  string `json:"hid"`
		Ldevice []string    `json:"ldevice"`
	} `json:"permission"`

}

