package event_test

import (
	"awesomeProject/event"
	"testing"
)

func TestDispatcher(t *testing.T) {
	called := false
	handler := func(ev interface{}) {
		if ev == 1 {
			called = true
		}
	}
	dispatcher := new(event.SimpleDispatcher)
	dispatcher.AddListener(event.Listener{EventName: "test", Handler: handler})
	dispatcher.Dispatch("test", 1)

	if !called {
		t.Fatal("Event listener was not called")
	}
}
