package mock

import "testing"

type DHT struct {
	Testing       *testing.T
	UpdateCalled  bool
	CelsiusCalled bool
}

func (m *DHT) Update() {
	m.UpdateCalled = true
}

func (m *DHT) Celsius() int16 {
	m.CelsiusCalled = true
	return 0
}
