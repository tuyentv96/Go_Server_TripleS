package api

import (
	modelctrl "../../controller/model"
	"encoding/json"
)

func HomeStatusRev(payload []byte)  (modelctrl.SendHomeStatus) {

	var m modelctrl.HomeStatus
	bytes := []byte(payload)

	err := json.Unmarshal(bytes, &m)
	ret := modelctrl.SendHomeStatus{Title: "MHSTATUS"}

	if err != nil {
		//fmt.Print("Error json")
		return modelctrl.SendHomeStatus{}

	}

	ret.Hid=m.Hid

	return ret
}