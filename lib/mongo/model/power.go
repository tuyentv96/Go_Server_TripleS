package model

import "time"

type DevicePower struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Type  int `json:"type" bson:"type"`
	Power int `json:"power" bson:"power"`
	Date  time.Time `json:"date" bson:"date"`
	Time  int64 `json:"time" bson:"time"`
}
