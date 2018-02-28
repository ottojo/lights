package lights

import (
	"github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

type Color struct {
	R float64
	G float64
	B float64
}

var mqttClient mqtt.Client

func init() {
	Reconnect()
}

func Reconnect() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://odroid.lan:1883")
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SetAll(c Color) {
	SetWindowLedStrip(c)
	SetMateLightAll(c)
}

const (
	MATELIGHT_RAINBOW = 'R'
	MATELIGHT_BALL    = 'B'
)

type MateLightAnimation rune

func SetMateLightAnimation(animation MateLightAnimation) {
	sendMqtt("lights/mateLight/pixel", string(animation))
}

func SetMateLightPixel(x, y int, c Color) {
	sendMqtt("/lights/mateLight/pixel", []byte{
		byte(x), byte(y),
		byte(c.R * 256), byte(c.G * 256), byte(c.B * 256)})
}

func SetMateLightAll(c Color) {
	sendMqtt("/lights/mateLight/all/red", strconv.Itoa(int(c.R*256)))
	sendMqtt("/lights/ledStripWindow/all/green", strconv.Itoa(int(c.G*256)))
	sendMqtt("/lights/ledStripWindow/all/blue", strconv.Itoa(int(c.B*256)))
}

func SetWindowLedStrip(c Color) {
	sendMqtt("/lights/ledStripWindow/red", strconv.Itoa(int(c.R*1024)))
	sendMqtt("/lights/ledStripWindow/green", strconv.Itoa(int(c.G*1024)))
	sendMqtt("/lights/ledStripWindow/blue", strconv.Itoa(int(c.B*1024)))

}

func sendMqtt(topic string, payload interface{}) {
	token := mqttClient.Publish(topic, 1, false, payload)
	token.Wait()
}
