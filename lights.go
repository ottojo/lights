package lights

// Values in the range [0,1]
type Color struct {
	R float64
	G float64
	B float64
}

func init() {
	Reconnect()
}

func SetAll(c Color) {
	SetWindowLedStrip(c)
	SetMateLightAll(c)
}
