package model

type Mcontrol struct {
	Uid string
	Did string
	Status int
}

type Mcontrolrespond struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
}

