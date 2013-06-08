package mpd

type Event interface {
	Subsystem() Subsys
}
