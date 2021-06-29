package util

import (
	"math/rand"
)

type Change struct {
	Width  int
	Height int
	State  bool
}

func ResetgridAndQueue(width, height int) (grid [][]bool, queue []Change) {
	queue = make([]Change, 0)

	grid = make([][]bool, width)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, height)
	}

	return grid, queue
}

func ConwayRules(grid [][]bool, queue []Change, x, y, maxX, maxY int) ([][]bool, []Change) {
	cc := grid[x][y]

	//get buurcellen
	var cells []bool
	if x != 0 {
		cells = append(cells, grid[x-1][y])
	}
	if y != 0 {
		cells = append(cells, grid[x][y-1])
	}
	if x != 0 && y != 0 {
		cells = append(cells, grid[x-1][y-1])
	}
	if x != maxX-1 {
		cells = append(cells, grid[x+1][y])
	}
	if y != maxY-1 {
		cells = append(cells, grid[x][y+1])
	}
	if x != 0 && y != maxY-1 {
		cells = append(cells, grid[x-1][y+1])
	}
	if x != maxX-1 && y != 0 {
		cells = append(cells, grid[x+1][y-1])
	}
	if x != maxX-1 && y != maxY-1 {
		cells = append(cells, grid[x+1][y+1])
	}

	// if w != 0 {
	// 	cells = append(cells, grid[w-1][h])
	// }
	//
	// if h != 0 {
	// 	cells = append(cells, grid[w][h-1])
	// }
	//
	// if w != 0 && h != 0 {
	// 	cells = append(cells, grid[w-1][h-1])
	// }
	//
	// if w != g.gameW-1 {
	// 	cells = append(cells, grid[w+1][h])
	// }
	//
	// if h != g.gameH-1 {
	// 	cells = append(cells, grid[w][h+1])
	// }
	//
	// if w != 0 && h != g.gameH-1 {
	// 	cells = append(cells, grid[w-1][h+1])
	// }
	//
	// if w != g.gameW-1 && h != 0 {
	// 	cells = append(cells, grid[w+1][h-1])
	// }
	//
	// if w != g.gameW-1 && h != g.gameH-1 {
	// 	cells = append(cells, grid[w+1][h+1])
	// }

	sum := 0
	for _, v := range cells {
		if v {
			sum += 1
		}
	}

	if cc && sum < 2 {
		queue = append(queue, Change{Width: x, Height: y, State: false})
	} else if cc && sum == 2 || cc && sum == 3 {
		queue = append(queue, Change{Width: x, Height: y, State: true})
	} else if cc && sum > 3 {
		queue = append(queue, Change{Width: x, Height: y, State: false})
	} else if !cc && sum == 3 {
		queue = append(queue, Change{Width: x, Height: y, State: true})
	}

	return grid, queue
}

func Randomgrid(grid [][]bool, queue []Change) ([][]bool, []Change) {
	for w := 0; w < len(grid); w++ {
		for h := 0; h < len(grid[w]); h++ {
			rBool := rand.Intn(2) == 1
			queue = append(queue, Change{Width: w, Height: h, State: rBool})
			grid[w][h] = rBool
		}
	}

	return grid, queue
}
