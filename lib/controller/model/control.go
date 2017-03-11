package model

type Control struct {
	Hid string
	Did string
	Status int
}

type Controlrsp struct {
	Title string `json:"title"`
	Rcode int `json:"rcode"`
}
