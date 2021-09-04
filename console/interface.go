package console

import (
	"io"
)

type Console interface {
	io.WriteCloser
}
