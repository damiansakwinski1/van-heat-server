package heating

type Device interface {
	Enable()
	Disable()
	IsEnabled() bool
}
