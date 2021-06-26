package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	winW     int
	winH     int
	gameW    int
	gameH    int
	grid     [][]bool
	queue    []change
	disabled bool
	gen      int
}

type change struct {
	width  int
	height int
	state  bool
}

var (
	cc  bool
	sum int
)

func (g *Game) Update() error {
	g.queue = nil
	if !g.disabled {
		grid := g.grid
		for w := 1; w < g.gameW-1; w++ {
			for h := 1; h < g.gameH-1; h++ {
				cc = grid[w][h]

				sum = 0
				for _, v := range []bool{
					grid[w+1][h],
					grid[w+1][h+1],
					grid[w][h+1],
					grid[w-1][h],
					grid[w-1][h-1],
					grid[w][h-1],
					grid[w+1][h-1],
					grid[w-1][h+1],
				} {
					if v {
						sum += 1
					}
				}

				if cc && sum < 2 {
					g.queue = append(g.queue, change{width: w, height: h, state: false})
				} else if cc && sum == 2 || cc && sum == 3 {
					g.queue = append(g.queue, change{width: w, height: h, state: true})
				} else if cc && sum > 3 {
					g.queue = append(g.queue, change{width: w, height: h, state: false})
				} else if !cc && sum == 3 {
					g.queue = append(g.queue, change{width: w, height: h, state: true})
				}
			}
		}
		g.gen += 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	ebiten.SetWindowTitle(fmt.Sprintf("Go Of Life | Generation: %v | FPS: %f", g.gen, ebiten.CurrentFPS()))

	for _, v := range g.queue {
		if v.state {
			g.grid[v.width][v.height] = true
			screen.Set(v.width, v.height, color.Black)
		} else {
			g.grid[v.width][v.height] = false
			screen.Set(v.width, v.height, color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameW, g.gameH
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	game := &Game{winW: 1440, winH: 900, queue: make([]change, 0)}
	game.gameW = game.winW / 4
	game.gameH = game.winH / 4
	game.grid = make([][]bool, game.gameW)

	ebiten.SetWindowSize(game.winW, game.winH)
	ebiten.SetWindowTitle(fmt.Sprintf("Go Of Life | FPS: %f", ebiten.CurrentFPS()))

	for i := 0; i < len(game.grid); i++ {
		game.grid[i] = make([]bool, game.gameH)
	}

	for w := 0; w < len(game.grid); w++ {
		for h := 0; h < len(game.grid[w]); h++ {
			game.grid[w][h] = rand.Intn(2) == 1
		}
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
