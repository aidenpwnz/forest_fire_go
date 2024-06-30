package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Forest struct {
	Tick   time.Duration
	Width  int
	Height int
	Cells  [][]Cell
}

func NewForest(width, height int, tick time.Duration) *Forest {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	// set a random cell on fire
	cells[rand.Intn(height)][rand.Intn(width)].State = Burning

	return &Forest{
		Tick:   tick,
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (f *Forest) Draw() {
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			fmt.Printf("%s", f.Cells[y][x].State.ToString())
		}
		fmt.Println()
	}
}

func (f *Forest) UpdateForest() {
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			neighbors := f.ConstructNeighbors(x, y)

			f.Cells[y][x].ShouldChangeState(neighbors)
		}
	}
}

type NeighborCells []*Cell

func (f *Forest) ConstructNeighbors(x, y int) NeighborCells {
	neighbors := make(NeighborCells, 4)

	up := f.constructCell(x, y-1)
	down := f.constructCell(x, y+1)
	left := f.constructCell(x-1, y)
	right := f.constructCell(x+1, y)

	neighbors = append(neighbors, up, down, left, right)

	return neighbors
}

func (f *Forest) constructCell(x, y int) *Cell {
	if x < 0 || x >= f.Width || y < 0 || y >= f.Height {
		return nil
	}
	return &f.Cells[y][x]
}

func (f *Forest) CanBurn() bool {
	canChange := 0
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			if f.Cells[y][x].CanChangeState(f.ConstructNeighbors(x, y)) {
				canChange++
			}
		}
	}

	return canChange == 0
}
