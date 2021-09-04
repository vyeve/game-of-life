package console

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

type console struct {
	sync.Mutex
	os     string
	writer io.Writer
	buf    bytes.Buffer
}

func New(wc io.WriteCloser) Console {
	if wc == nil {
		wc = os.Stderr
	}
	c := &console{
		writer: wc,
		os:     runtime.GOOS,
	}
	return c
}

func (c *console) WriteFrame(p []byte) (err error) {
	c.Lock()
	defer c.Unlock()
	_, err = c.buf.Write(p)
	if err != nil {
		return err
	}
	return c.flush()
}

func (c *console) Clear() error {
	c.Lock()
	defer c.Unlock()
	c.buf.Reset()
	return c.clearTerm()
}

func (c *console) flush() error {
	err := c.clearTerm()
	if err != nil {
		return err
	}
	if c.buf.Len() == 0 {
		return nil
	}
	_, err = c.writer.Write(c.buf.Bytes())
	c.buf.Reset()
	return err
}

func (c *console) clearTerm() error {
	var cmd *exec.Cmd
	if c.os != "windows" {
		cmd = exec.Command("clear")
	} else {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = c.writer
	return cmd.Run()
}
