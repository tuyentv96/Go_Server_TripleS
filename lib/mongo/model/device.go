package model


	// User represents the structure of our resource
type Device struct {
		Did     string `json:"did" bson:"did"`
		Hid     string `json:"hid" bson:"hid"`
		Dname   string `json:"dname" bson:"dname"`
		Status  int `json:"status" bson:"status"`
		Type  int `json:"type" bson:"type"`
		Roomid  string `json:"roomid" bson:"roomid"`
		Roomname  string `json:"roomname" bson:"roomname"`
	}

type LDevice struct {
	UID   string `json:"uid"`
	Uname string `json:"uname"`
	Lhome []struct {
		Hid  string `json:"hid"`
		Type int    `json:"type"`
	} `json:"lhome"`
	Ldevice []Device
	Permission []struct {
		Hid  string `json:"hid"`
		Did string    `json:"did"`
	} `json:"permission"`
}

type LHomeDevice struct {
	Ldevice []Device `json:"ldevice"`
}

type HistoryDevice struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Uid  string `json:"uid" bson:"uid"`
	Time int64 `json:"time" bson:"time"`
}


