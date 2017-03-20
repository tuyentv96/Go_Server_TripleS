package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	model "../controller/model"
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

func MqttPublishSingle(topic string,qos byte,retain bool,payload []byte)  {
	opts := MQTT.NewClientOptions().AddBroker(Mqttbroker)
	opts.SetClientID(RandStringRunes(12))
	opts.SetCleanSession(true)


	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token:= c.Publish(topic,qos,retain,payload)
	token.Wait()
	c.Disconnect(0)
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


