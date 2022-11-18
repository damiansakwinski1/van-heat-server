package event

type Dispatcher interface {
	Dispatch(eventName string, event interface{})
	AddListener(l Listener)
}

type SimpleDispatcher struct {
	listeners []Listener
}

func (d *SimpleDispatcher) Dispatch(eventName string, event interface{}) {
	for _, listener := range d.listeners {
		if listener.EventName != eventName {
			continue
		}

		listener.Handler(event)
	}
}

func (d *SimpleDispatcher) AddListener(l Listener) {
	d.listeners = append(d.listeners, l)
}
