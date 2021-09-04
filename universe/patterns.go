// Source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
//
package universe

import (
	"math/rand"
	"time"
)

const (
	GlidePattern = iota + 1
	PulsePattern
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
	gliderData = [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
	}
	/*
		| | |█| | |
		| | | |█| |
		| |█|█|█| |
		| | | | | |
	*/
	pulsarData = [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, false, false, true, true, true, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, true, true, true, false, false, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, true, false, false, false, false, true, false, true, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	}
	/*
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
	*/
	mwssData = [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, true, false, false, false},
		{false, false, true, true, true, false, true, true, false, false},
		{false, false, true, true, true, true, true, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
	/*
		| | | | | | | | | | |
		| | | | | | | | | | |
		| | | | | |█|█| | | |
		| | |█|█|█| |█|█| | |
		| | |█|█|█|█|█| | | |
		| | | |█|█|█| | | | |
		| | | | | | | | | | |
		| | | | | | | | | | |
	*/
	gunData = [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false},
		{false, true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	}
	/*
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
	*/
	decathlonData = [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
	}
	/*
		| | | | | | | | | | | |
		| | | | | |█| | | | | |
		| | | | | |█| | | | | |
		| | | | |█|█|█| | | | |
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
	*/
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
	return newPattern(gliderData)
}

func NewPulsePattern() Pattern {
	return newPattern(pulsarData)
}

func NewMWSSPattern() Pattern {
	return newPattern(mwssData)
}

func NewGunPattern() Pattern {
	return newPattern(gunData)
}

func NewDecathlonPattern() Pattern {
	return newPattern(decathlonData)
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
