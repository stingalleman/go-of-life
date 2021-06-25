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

func (g *Game) Update() error {
	g.queue = nil
	if !g.disabled {
		grid := g.grid
		for w := 1; w < g.winW-1; w++ {
			for h := 1; h < g.winH-1; h++ {

				// if w > 1 && w < g.winW-1 {
				// 	if h > 1 && h < g.winH-1 {
				cc := grid[w][h]

				sum := 0
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
	return g.winW, g.winH
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetWindowSize(2560, 1080)
	ebiten.SetWindowTitle(fmt.Sprintf("Go Of Life | FPS: %f", ebiten.CurrentFPS()))

	game := &Game{winW: 2560, winH: 1080, disabled: false, grid: make([][]bool, 2560), queue: make([]change, 0)}

	for i := 0; i < len(game.grid); i++ {
		game.grid[i] = make([]bool, 1080)
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
