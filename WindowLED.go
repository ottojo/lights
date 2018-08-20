package lights

import (
	"strconv"
)

func SetWindowLedStrip(c Color) {
	sendMqtt("/lights/ledStripWindow/red", strconv.Itoa(int(c.R*1023)))
	sendMqtt("/lights/ledStripWindow/green", strconv.Itoa(int(c.G*1023)))
	sendMqtt("/lights/ledStripWindow/blue", strconv.Itoa(int(c.B*1023)))

}
