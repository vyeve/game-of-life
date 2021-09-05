package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/vyeve/game-of-life/console"
	"github.com/vyeve/game-of-life/universe"
)

const (
	numLoops = 100
)

func main() {
	var (
		p   int
		pat universe.Pattern
	)
	flag.IntVar(&p, "p", 5, "init pattern to run")
	flag.Parse()

	switch p {
	case universe.GlidePattern:
		pat = universe.NewGliderPattern()
	case universe.PulsePattern:
		pat = universe.NewPulsePattern()
	case universe.MWSSPattern:
		pat = universe.NewMWSSPattern()
	case universe.GunPattern:
		pat = universe.NewGunPattern()
	case universe.DecathlonPattern:
		pat = universe.NewDecathlonPattern()
	case universe.RandomPattern:
		pat = universe.NewRandomPattern()
	default:
		pat = universe.NewPulsePattern()
	}

	wr := console.New(os.Stderr)
	u := universe.NewUniverse(pat)

	for i := 0; i < numLoops; i++ {
		p := u.State()
		if err := wr.WriteFrame(p); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 500)
		u = u.NextGen()
	}

	if err := wr.Clear(); err != nil {
		log.Fatal(err)
	}
}
