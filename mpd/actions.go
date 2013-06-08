package mpd

type Action interface {
	
}

const (
	ActionPrev int = iota
	ActionNext
	ActionPlay
	ActionPause
	ActionToggle
)
