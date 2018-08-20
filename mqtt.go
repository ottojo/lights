package lights

import "github.com/eclipse/paho.mqtt.golang"

var mqttClient mqtt.Client

func Reconnect() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://odroid.lan:1883")
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func sendMqtt(topic string, payload interface{}) {
	token := mqttClient.Publish(topic, 1, false, payload)
	token.Wait()
}
