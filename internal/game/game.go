package game

import (
	c "minesweeper/internal/cell"
)

type Status int

const (
	StatusInProgress = iota
	StatusWon
	StatusLost
)

type Game struct {
	Grid          [][]*c.Cell
	Status        Status
	Height, Width int
	BlackHoleNum  int
}

func NewGame(grid [][]*c.Cell, status Status, height, width, blackHoleNum int) *Game {
	return &Game{
		Grid:         grid,
		Status:       status,
		Height:       height,
		Width:        width,
		BlackHoleNum: blackHoleNum,
	}
}
