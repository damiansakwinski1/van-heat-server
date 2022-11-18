package sensor

type OneWire struct{}

func (o OneWire) Celsius() int16 {
	return 0
}

func (o OneWire) Update() {
}
