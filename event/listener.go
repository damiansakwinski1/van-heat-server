package event

type Handler func(event interface{})

type Listener struct {
	EventName string
	Handler   Handler
}
