package console

type Console interface {
	// WriteFrame clears console and writes bytes
	WriteFrame(frame []byte) error
	// Clear clears console
	Clear() error
}
