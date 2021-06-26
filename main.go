package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	sleep    int64
	started  bool
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
	// handle user input
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.disabled {
			g.disabled = false
		} else {
			g.disabled = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.sleep += 100
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.sleep -= 100
	}

	if !g.disabled {
		g.queue = nil
		grid := g.grid
		for w := 0; w < g.gameW; w++ {
			for h := 0; h < g.gameH; h++ {
				g.started = true
				cc = grid[w][h]

				//get buurcellen
				var cells []bool
				if w != 0 {
					cells = append(cells, grid[w-1][h])
				}
				if h != 0 {
					cells = append(cells, grid[w][h-1])
				}
				if w != 0 && h != 0 {
					cells = append(cells, grid[w-1][h-1])
				}
				if w != g.gameW-1 {
					cells = append(cells, grid[w+1][h])
				}
				if h != g.gameH-1 {
					cells = append(cells, grid[w][h+1])
				}
				if w != 0 && h != g.gameH-1 {
					cells = append(cells, grid[w-1][h+1])
				}
				if w != g.gameW-1 && h != 0 {
					cells = append(cells, grid[w+1][h-1])
				}
				if w != g.gameW-1 && h != g.gameH-1 {
					cells = append(cells, grid[w+1][h+1])
				}

				sum = 0
				for _, v := range cells {
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
	var paused string
	if g.disabled {
		paused = "| Paused "
	}
	title := fmt.Sprintf("Go Of Life | Generation: %v | Delay: %v ms %s| FPS: %f", g.gen, g.sleep, paused, ebiten.CurrentFPS())
	ebiten.SetWindowTitle(title)
	screen.Fill(color.White)

	for _, v := range g.queue {
		if v.state {
			g.grid[v.width][v.height] = true
			screen.Set(v.width, v.height, color.Black)
		} else {
			g.grid[v.width][v.height] = false
			screen.Set(v.width, v.height, color.White)
		}
	}
	time.Sleep(time.Millisecond * time.Duration(g.sleep))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameW, g.gameH
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	game := &Game{winW: 720, winH: 450, disabled: true, queue: make([]change, 0)}
	game.gameH = game.winH / 6
	game.gameW = game.winW / 6
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
