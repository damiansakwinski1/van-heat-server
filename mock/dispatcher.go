package mock

import "awesomeProject/event"

type Dispatcher struct {
	AddListenerCalled bool
	DispatchEventName string
	DispatchEvent     interface{}
	DispatchCalled    bool
}

func (d *Dispatcher) AddListener(l event.Listener) {
	d.AddListenerCalled = true
}

func (d *Dispatcher) Dispatch(eventName string, event interface{}) {
	d.DispatchEventName = eventName
	d.DispatchEvent = event
	d.DispatchCalled = true
}
