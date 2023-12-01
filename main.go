package main

import (
	"fmt"
	"time"
)

const (
	maxRows         = 25
	maxCols         = 25
	noOfGenerations = 25 // Number of times the universe will regenrated
)

// Cell represents a single cell in the Game of Life.
type Cell struct {
	alive bool
}

// Universe represents the game board.
type Universe struct {
	grid       [][]Cell
	rows, cols int
}

// NewUniverse creates a new Universe with the specified dimensions.
func NewUniverse(rows, cols int) *Universe {
	grid := make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
	}
	return &Universe{grid: grid, rows: rows, cols: cols}
}

// SetCell sets the state of a cell in the universe.
func (u *Universe) SetCell(row, col int, alive bool) {
	// Condition to handle outofbound errors
	if row < maxRows && col < maxCols {
		u.grid[row][col].alive = alive
	}
}

// Print prints the current state of the universe to the console.
func (u *Universe) Print() {
	for _, row := range u.grid {
		for _, cell := range row {
			if cell.alive {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// NextGeneration computes the next generation of the universe based on the rules.
func (u *Universe) NextGeneration() {
	nextGen := NewUniverse(u.rows, u.cols)

	for i := 0; i < u.rows; i++ {
		for j := 0; j < u.cols; j++ {
			liveNeighbors := u.countLiveNeighbors(i, j)
			currentCell := u.grid[i][j]

			// Apply the rules
			if currentCell.alive && (liveNeighbors < 2 || liveNeighbors > 3) {
				nextGen.SetCell(i, j, false) // Rule 1 and 3
			} else if currentCell.alive && (liveNeighbors == 2 || liveNeighbors == 3) {
				nextGen.SetCell(i, j, true) // Rule 2
			} else if !currentCell.alive && liveNeighbors == 3 {
				nextGen.SetCell(i, j, true) // Rule 4
			}
		}
	}

	u.grid = nextGen.grid
}

// countLiveNeighbors counts the number of live neighbors for a given cell.
func (u *Universe) countLiveNeighbors(row, col int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < u.rows && j >= 0 && j < u.cols && !(i == row && j == col) {
				if u.grid[i][j].alive {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// Create a  universe with max rows and cols defined in constants
	universe := NewUniverse(maxRows, maxCols)

	// Set the initial state with the 'Glider' pattern and additional cells
	gliderPattern := [5][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
	for _, coord := range gliderPattern {
		universe.SetCell(coord[0], coord[1], true)
	}

	// Run for no of generations and print the state after each generation
	for generation := 1; generation <= noOfGenerations; generation++ {
		fmt.Printf("Generation %d:\n", generation)
		universe.Print()
		time.Sleep(time.Second) // Pause for visibility
		universe.NextGeneration()
	}
}
