package event

type Type int

const (
	Login = iota
	ChannelJoin
	Command
	Unknown
)

type Event struct {
	Type Type
	Data map[string]string
}
