package temperature_test

import (
	"awesomeProject/mock"
	"awesomeProject/sensor"
	"awesomeProject/temperature"
	"testing"
)

func TestDHT(t *testing.T) {
	dht := new(sensor.DHT)
	dht.Celsius()
}

func TestWaterUpdate(t *testing.T) {
	mockDht := mock.DHT{Testing: t}
	mockDisp := mock.Dispatcher{}
	water := temperature.Water{Sensor: &mockDht, Dispatcher: &mockDisp}
	water.Update()

	if !mockDht.UpdateCalled {
		t.Fatal("Expected a call to DHT.Update()")
	}

	if !mockDisp.DispatchCalled {
		t.Fatal("Expected a call to Dispatcher.Dispatch()")
	}

	expectedEventName := "water.temperature.update.after"
	if mockDisp.DispatchEventName != expectedEventName {
		t.Fatalf("Got invalid Dispatcher.Dispatch event name. Expected: %s, got: %s", "water.temperature.update.after", mockDisp.DispatchEventName)
	}

	if mockDisp.DispatchEvent != 1 {
		t.Fatalf("Got invalid Dispatcher.Dispatch event. Expected: %d, got: %d", 1, mockDisp.DispatchEvent)
	}
}
