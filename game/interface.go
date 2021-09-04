package game

type Universe interface {
	NextGen() Universe
	State() []byte
}

type Pattern interface {
	Data() [][]bool
	Width() int
	Height() int
}
