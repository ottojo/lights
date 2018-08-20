package lights

import "strconv"

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
		byte(c.R * 255), byte(c.G * 255), byte(c.B * 255)})
}

func SetMateLightAll(c Color) {
	sendMqtt("/lights/mateLight/all/red", strconv.Itoa(int(c.R*255)))
	sendMqtt("/lights/mateLight/all/green", strconv.Itoa(int(c.G*255)))
	sendMqtt("/lights/mateLight/all/blue", strconv.Itoa(int(c.B*255)))
}
