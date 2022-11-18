package heating

type Webasto struct{}

func (w Webasto) Enable() {}

func (w Webasto) Disable() {}

func (w Webasto) IsEnabled() bool {
	return false
}
