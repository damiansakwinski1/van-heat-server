package heating

type Fan struct{}

func (f Fan) Enable()  {}
func (f Fan) Disable() {}
func (f Fan) IsEnabled() bool {
	return false
}
func (f Fan) SetSpeed() {}
