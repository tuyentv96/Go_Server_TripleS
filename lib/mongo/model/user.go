package model

type (
	// User represents the structure of our resource
	User struct {
		Uid     string `json:"uid" bson:"uid"  form:"uid"`
		Pwd     string `json:"pwd" bson:"pwd"  form:"pwd"`
		Uname   string `json:"uname" bson:"uname"  form:"uname"`

	}
)
