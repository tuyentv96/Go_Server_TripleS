package model

type User struct {
		Uid     string `json:"uid" bson:"uid"  form:"uid"`
		Pwd     string `json:"pwd" bson:"pwd"  form:"pwd"`
		Uname   string `json:"uname" bson:"uname"  form:"uname"`

	}

type Userpsmdevice struct {
	Devices []struct {
		Device []string `json:"device"`
		Hid    string   `json:"hid"`
	} `json:"devices"`
	Types []struct {
		Hid  string `json:"hid"`
		Type int    `json:"type"`
	} `json:"type"`
	UID   string `json:"uid"`
	Uname string `json:"uname"`
}