// Patterns source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
//
package universe

import (
	"bufio"
	"math/rand"
	"strings"
	"time"
)

const (
	GliderPattern = iota + 1
	PulsarPattern
	MWSSPattern
	GunPattern
	DecathlonPattern
	RandomPattern
)

const (
	defaultWidth  = 40
	defaultHeight = 20
)

var (
	glider = `
	| | |█| | |
	| | | |█| |
	| |█|█|█| |
	| | | | | |	
	`
	pulsar = `
	| | | | | | | | | | | | | | | | | |
	| | | | | | | | | | | | | | | | | |
	| | | | |█|█|█| | | |█|█|█| | | | |
	| | | | | | | | | | | | | | | | | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | | | |█|█|█| | | |█|█|█| | | | |
	| | | | | | | | | | | | | | | | | |
	| | | | |█|█|█| | | |█|█|█| | | | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | |█| | | | |█| |█| | | | |█| | |
	| | | | | | | | | | | | | | | | | |
	| | | | |█|█|█| | | |█|█|█| | | | |
	| | | | | | | | | | | | | | | | | |
	| | | | | | | | | | | | | | | | | |	
	`
	mwss = `
	| | | | | | | | | | |
	| | | | | | | | | | |
	| | | | | |█|█| | | |
	| | |█|█|█| |█|█| | |
	| | |█|█|█|█|█| | | |
	| | | |█|█|█| | | | |
	| | | | | | | | | | |
	| | | | | | | | | | |	
	`
	gun = `
	| | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |
	| | | | | | | | | | | | | | | | | | | | | | | | | | |█| | | | | | | | | | | |
	| | | | | | | | | | | | | | | | | | | | | | | | |█| |█| | | | | | | | | | | |
	| | | | | | | | | | | | | | |█|█| | | | | | |█|█| | | | | | | | | | | |█|█| |
	| | | | | | | | | | | | | |█| | | |█| | | | |█|█| | | | | | | | | | | |█|█| |
	| |█|█| | | | | | | | | |█| | | | | |█| | | |█|█| | | | | | | | | | | | | | |
	| |█|█| | | | | | | | | |█| | | |█| |█|█| | | | |█| |█| | | | | | | | | | | |
	| | | | | | | | | | | | |█| | | | | |█| | | | | | | |█| | | | | | | | | | | |
	| | | | | | | | | | | | | |█| | | |█| | | | | | | | | | | | | | | | | | | | |
	| | | | | | | | | | | | | | |█|█| | | | | | | | | | | | | | | | | | | | | | |
	| | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |	
	`
	decathlon = `
	| | | | | | | | | | | |
	| | | | | |█| | | | | |
	| | | | | |█| | | | | |
	| | | | |█|█|█| | | | |
	| | | | | | | | | | | |
	| | | | | | | | | | | |
	| | | | |█|█|█| | | | |
	| | | | | |█| | | | | |
	| | | | | |█| | | | | |
	| | | | | |█| | | | | |
	| | | | | |█| | | | | |
	| | | | |█|█|█| | | | |
	| | | | | | | | | | | |
	| | | | | | | | | | | |
	| | | | |█|█|█| | | | |
	| | | | | |█| | | | | |
	| | | | | |█| | | | | |
	| | | | | | | | | | | |	
	`
)

type (
	pattern struct {
		data [][]bool
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newPattern(data [][]bool) Pattern {
	return pattern{
		data: data,
	}
}

func NewGliderPattern() Pattern {
	return newPattern(parseText(glider))
}

func NewPulsarPattern() Pattern {
	return newPattern(parseText(pulsar))
}

func NewMWSSPattern() Pattern {
	return newPattern(parseText(mwss))
}

func NewGunPattern() Pattern {
	return newPattern(parseText(gun))
}

func NewDecathlonPattern() Pattern {
	return newPattern(parseText(decathlon))
}

func NewRandomPattern() Pattern {
	return newPattern(randomizeSlice(defaultWidth, defaultHeight))
}

func (p pattern) Data() [][]bool {
	return p.data
}

func (p pattern) Width() int {
	return len(p.data[0])
}

func (p pattern) Height() int {
	return len(p.data)
}

func randomizeSlice(w, h int) [][]bool {
	data := make([][]bool, h)
	for i := range data {
		data[i] = make([]bool, w)
	}
	for i := 0; i < w*h/5; i++ {
		data[rand.Intn(h)][rand.Intn(w)] = true
	}
	return data
}

func parseText(src string) [][]bool {
	scanner := bufio.NewScanner(strings.NewReader(src))
	out := make([][]bool, 0)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "|")
		row := make([]bool, 0, len(line))
		for _, c := range line {
			if len(c) == 0 {
				continue
			}
			var isAlive bool
			if c == string(aliveChar) {
				isAlive = true
			}
			row = append(row, isAlive)
		}
		if len(row) > 0 {
			out = append(out, row)
		}
	}
	return out
}
