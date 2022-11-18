package heating

type Pump struct{}

func (p Pump) Enable() {}

func (p Pump) Disable() {}

func (p Pump) IsEnabled() bool {
	return false
}
