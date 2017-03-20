package api

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"

)

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
