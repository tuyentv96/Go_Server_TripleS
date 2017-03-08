package model

type Mgethome struct {
	Uid string `json:"uid"`
}

type Mgethomerespond struct {
	Title string `json:"title"`
	Rcode int    `json:"rcode"`
	Lhome []struct {
		Hid   string `json:"hid"`
		Hname string `json:"hname"`
	} `json:"lhome"`

}

