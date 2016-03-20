package event

type Type int

const (
	Login = iota
	Unknown
)

type Event struct {
	Type Type
}
