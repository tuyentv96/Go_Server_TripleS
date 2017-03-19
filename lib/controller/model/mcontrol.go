package model

type Mcontrol struct {
	Uid string
	Hid string
	Did string
	Status int
}

type Mcontrolrespond struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
	Did string `json:"did"`
	Hid string `json:"hid"`
}

