package temperature

import (
	"awesomeProject/event"
)

type Sensor interface {
	Celsius() int16
	Update()
}

type Water struct {
	Sensor
	event.Dispatcher
}

func (w *Water) Update() {
	w.Sensor.Update()
	w.Dispatcher.Dispatch("water.temperature.update.after", 1)
}

type Ambient struct {
	Sensor
	event.Dispatcher
}

func (a *Ambient) Update() {
	a.Sensor.Update()
	a.Dispatcher.Dispatch("ambient.temperature.update", 1)
}
