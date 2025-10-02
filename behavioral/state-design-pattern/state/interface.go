package state

type State interface {
	Publish()
	GetState() string
}
