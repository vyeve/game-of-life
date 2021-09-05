package universe

import (
	"bytes"
	"runtime"
	"strings"
)

const (
	deadChar  = ' '
	aliveChar = '\u2588'
)

var (
	colorize, reset string
	topBorder       = "_"
	bottomBorder    = "_"
)

// unix specific symbols
func init() {
	if runtime.GOOS != "windows" {
		colorize = "\033[33m"
		reset = "\033[0m"
		topBorder = "\u23bd"
		bottomBorder = "\u23ba"
	}
}

type universe struct {
	fields [][]bool
	width  int
	height int
}

func NewUniverse(p Pattern) Universe {
	u := newUniverse(p.Width(), p.Height())
	for i, row := range p.Data() {
		for j, v := range row {
			u.set(j, i, v)
		}
	}
	return u
}

func NewGliderUniverse() Universe {
	return NewUniverse(NewGliderPattern())
}

func NewPulsarUniverse() Universe {
	return NewUniverse(NewPulsarPattern())
}

func NewMWSSUniverse() Universe {
	return NewUniverse(NewMWSSPattern())
}

func NewGunUniverse() Universe {
	return NewUniverse(NewGunPattern())
}

func NewDecathlonUniverse() Universe {
	return NewUniverse(NewDecathlonPattern())
}

func NewRandomUniverse() Universe {
	return NewUniverse(NewRandomPattern())
}

func newUniverse(width, height int) *universe {
	u := &universe{
		width:  width,
		height: height,
		fields: make([][]bool, height),
	}
	for i := range u.fields {
		u.fields[i] = make([]bool, u.width)
	}
	return u
}

func (u *universe) isAlive(w, h int) bool {
	w += u.width
	w %= u.width
	h += u.height
	h %= u.height
	return u.fields[h][w]
}

func (u *universe) set(w, h int, lifeness bool) {
	u.fields[h][w] = lifeness
}

func (u *universe) next(w, h int) bool {
	neighbors := 0
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if (j != 1 || i != 1) && // skip itself
				u.isAlive(w+i-1, h+j-1) {
				neighbors++
			}
		}
	}
	if u.isAlive(w, h) {
		return neighbors == 2 || neighbors == 3
	}
	return neighbors == 3
}

func (u *universe) NextGen() Universe {
	gen := newUniverse(u.width, u.height)
	for h, row := range u.fields {
		for w := range row {
			gen.set(w, h, u.next(w, h))
		}
	}
	return gen
}

func (u *universe) State() []byte {
	var buf bytes.Buffer
	var n int
	for i, row := range u.fields {
		buf.WriteRune('\t')
		if i == 0 {
			n = len(row)
			buf.WriteRune(' ')
			writeTopBorder(&buf, n)
			buf.WriteRune('\n')
			buf.WriteRune('\t')
		}
		buf.WriteRune('|')
		buf.WriteString(colorize)
		for _, val := range row {
			if val {
				buf.WriteRune(aliveChar)
			} else {
				buf.WriteRune(deadChar)
			}
		}
		buf.WriteString(reset)
		buf.WriteRune('|')
		buf.WriteRune('\n')
	}
	buf.WriteRune('\t')
	buf.WriteRune(' ')
	writeBottomBorder(&buf, n)
	buf.WriteRune(' ')
	buf.WriteRune('\n')
	return buf.Bytes()
}

func writeTopBorder(buf *bytes.Buffer, n int) {
	buf.WriteString(strings.Repeat(topBorder, n))
}

func writeBottomBorder(buf *bytes.Buffer, n int) {
	buf.WriteString(strings.Repeat(bottomBorder, n))
}
