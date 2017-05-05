package model


type UpdatePower struct {
	Hid string `json:"hid"`
	Did string `json:"did"`
	Power int `json:"power"`
}



type UpdatePowerRsp struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
}

type MSyncPower struct {
	Title string `json:"title"`
	Hid string `json:"hid"`
	Did string `json:"did"`
	Power int `json:"power"`
}
