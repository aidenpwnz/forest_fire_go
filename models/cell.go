package models

import "math/rand"

type CellInterface interface {
	New() *Cell
	ShouldStartBurning()
}

type Cell struct {
	State
}

func NewCell() *Cell {
	return &Cell{State: Tree}
}

func (c *Cell) ShouldChangeState(neighbors NeighborCells) {
	r := rand.Float64()

	// If any of the up, down, left, or right neighbors are burning, the cell will also have a chance start burning.
	if c.State == Tree {
		for _, neighbor := range neighbors {
			if neighbor != nil && neighbor.State == Burning && r < 0.8 {
				c.State = Burning
				break
			}
		}
	} else if c.State == Burning {
		c.State = Burnt
	} else if c.State == Burnt {
		c.State = Empty
	}
}

func (c *Cell) CanChangeState(neighbors NeighborCells) bool {
	for _, neighbor := range neighbors {
		if neighbor != nil && neighbor.State == Burning {
			return true
		}
	}

	return false
}
