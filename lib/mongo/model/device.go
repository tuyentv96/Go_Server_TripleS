package model

type (
	// User represents the structure of our resource
	Device struct {
		Did     string `json:"did" bson:"did"`
		Dname   string `json:"dname" bson:"dname"`
		Status  string `json:"status" bson:"status"`
		Type  string `json:"type" bson:"type"`
		Roomid  string `json:"roomid" bson:"roomid"`
		Roomname  string `json:"roomname" bson:"roomname"`
	}
)
