package api

import (
	model "../../controller/model"
	"strings"

	"math/rand"

)

const	(
	Mqttbroker="tcp://trantuyen.name.vn:1883"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}



func CutTopic(topic string)  (info model.RqDetail,b bool){
	index:= strings.Index(topic,"/")
	if  index>-1 {
		info.Topic=topic[index+1:]
		info.Cid=topic[:index]
		return info,false
	}
	return info,true

}
