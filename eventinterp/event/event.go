package event

type Type int

const (
	Login = iota
	ChannelJoin
	Unknown
)

type Event struct {
	Type Type
	Data map[string]string
}
