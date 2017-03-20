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
	Status int `json:"status"`
}

type Msync struct {
	Title string `json:"title"`
	Did string `json:"did"`
	Hid string `json:"hid"`
	Status int `json:"status"`
}
