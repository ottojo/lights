package lights

import (
	"github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

var mqttClient mqtt.Client

func Reconnect() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://odroid.lan:1883")
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SetWindowLedStrip(c Color) {
	sendMqtt("/lights/ledStripWindow/red", strconv.Itoa(int(c.R*1023)))
	sendMqtt("/lights/ledStripWindow/green", strconv.Itoa(int(c.G*1023)))
	sendMqtt("/lights/ledStripWindow/blue", strconv.Itoa(int(c.B*1023)))

}

func sendMqtt(topic string, payload interface{}) {
	token := mqttClient.Publish(topic, 1, false, payload)
	token.Wait()
}
