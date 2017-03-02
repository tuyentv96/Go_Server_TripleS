package model

type (
	// User represents the structure of our resource
	Device struct {
		Id     string `json:"id" bson:"id"  form:"id"`
		Name   string `json:"name" bson:"name"  form:"name"`
		Email  string `json:"email" bson:"email"  form:"email"`

	}
)
