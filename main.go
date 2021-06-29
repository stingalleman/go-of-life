package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stingalleman/go-of-life/game"
	"github.com/stingalleman/go-of-life/util"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	g := &game.Game{WinW: 720, WinH: 450, Disabled: true, Started: false}
	g.GameH = g.WinH / 6
	g.GameW = g.WinW / 6

	g.Grid, g.Queue = util.ResetgridAndQueue(g.GameW, g.GameH)

	ebiten.SetWindowSize(g.WinW, g.WinH)
	ebiten.SetWindowTitle(fmt.Sprintf("Go Of Life | Generation: %v | Delay: %v ms | FPS: %f", g.Gen, g.Sleep, ebiten.CurrentFPS()))

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
