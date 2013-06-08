package mpd

type Subsys int

const (
	SubsystemDatabase Subsys = iota
	SubsystemUpdate
	SubsystemStoredPlaylist
	SubsystemPlaylist
	SubsystemPlayer
	SubsystemMixer
	SubsystemOutput
	SubsystemOptions
	SubsystemSticker
	SubsystemSubscription
	SubsystemMessage
)
