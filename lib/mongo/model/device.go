package model

type (
	// User represents the structure of our resource
	Device struct {
		Did     string `json:"did" bson:"did"`
		Hid     string `json:"hid" bson:"hid"`
		Dname   string `json:"dname" bson:"dname"`
		Status  int `json:"status" bson:"status"`
		Type  int `json:"type" bson:"type"`
		Roomid  string `json:"roomid" bson:"roomid"`
		Roomname  string `json:"roomname" bson:"roomname"`
	}
)
