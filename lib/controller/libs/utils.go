package libs

import "strings"

func CutTopic(topic string)  (info model.RqDetail,b bool){
	index:= strings.Index(topic,"/")
	if  index>-1 {
		info.Topic=topic[index+1:]
		info.Cid=topic[:index]
		return info,false
	}
	return info,true

}