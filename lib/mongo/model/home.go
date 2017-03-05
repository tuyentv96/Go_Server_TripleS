package model

type (
	// User represents the structure of our resource
	Home struct {
		Hid     string `json:"hid" bson:"hid" `
		Hname   string `json:"hname" bson:"hname" `
		Address  string `json:"address" bson:"address" `
		Rooms  []string `json:"rooms" `

	}
)