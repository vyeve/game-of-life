package console

import (
	"io"
)

type Console interface {
	WriteFrame(frame []byte) error
	io.Closer
}
