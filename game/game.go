package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/stingalleman/go-of-life/util"
)

type Game struct {
	WinW     int
	WinH     int
	GameW    int
	GameH    int
	Grid     [][]bool
	Queue    []util.Change
	Disabled bool
	Gen      int
	Sleep    int64
	Started  bool
}

var (
	cc  bool
	sum int
)

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= 0 && x < g.GameW && y >= 0 && y < g.GameH {
			g.Queue = append(g.Queue, util.Change{Width: x, Height: y, State: true})
		}
	}

	// handle user input
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.Disabled {
			g.Disabled = false
		} else {
			g.Disabled = true
		}

		if !g.Started {
			g.Started = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Grid, g.Queue = util.ResetgridAndQueue(g.GameW, g.GameH)
		g.Grid, g.Queue = util.Randomgrid(g.Grid, g.Queue)
		g.Gen = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.Sleep += 100
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.Sleep -= 100
	}

	if !g.Disabled {
		g.Queue = nil
		for w := 0; w < g.GameW; w++ {
			for h := 0; h < g.GameH; h++ {
				g.Started = true

				g.Grid, g.Queue = util.ConwayRules(g.Grid, g.Queue, w, h, g.GameW, g.GameH)
			}

		}
		g.Gen += 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var paused string
	if g.Disabled {
		paused = "| Paused "
	}
	title := fmt.Sprintf("Go Of Life | Generation: %v | Delay: %v ms %s| FPS: %f", g.Gen, g.Sleep, paused, ebiten.CurrentFPS())
	ebiten.SetWindowTitle(title)
	screen.Fill(color.White)

	for _, v := range g.Queue {
		if v.State {
			g.Grid[v.Width][v.Height] = true
			screen.Set(v.Width, v.Height, color.Black)
		} else {
			g.Grid[v.Width][v.Height] = false
			screen.Set(v.Width, v.Height, color.White)
		}
	}
	time.Sleep(time.Millisecond * time.Duration(g.Sleep))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.GameW, g.GameH
}
