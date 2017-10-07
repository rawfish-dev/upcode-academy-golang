package interfaces

type DataProvider interface {
	BuildMessagePad() string
	ReadAll() []string
}
