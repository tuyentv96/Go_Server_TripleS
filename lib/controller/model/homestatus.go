package model

type HomeStatus struct {

	Hid string `json:"hid"`
}

type SendHomeStatus struct {
	Title string `json:"title"`
	Hid string `json:"hid"`
	Status int `json:"status"`
}