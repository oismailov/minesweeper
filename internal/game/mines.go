package game

import (
	"math/rand"
	c "minesweeper/internal/cell"
	"time"
)

func (g *Game) plantBlackHoles() {
	blackHoleLocations := g.getBlackHoleLocations()

	//plant black holes into cells
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			cellNumber := i*g.Height + j
			//check if randomly shuffled black holes slice has desired cell number
			//if yes - place black hole into that cell
			if blackHoleLocations[cellNumber] {
				lackHoleCell := g.Grid[i][j]
				lackHoleCell.Type = c.TypeBlackHole
				for _, adjacentCells := range lackHoleCell.GetAdjacentCells(g.Grid) {
					adjacentCells.AdjacentBlackHoles++
				}
			}
		}
	}
}

func (g *Game) getBlackHoleLocations() []bool {
	//build a slice with size of game Grid Height*Width and set false to all it's values
	blackHoleLocations := make([]bool, g.Height*g.Width)
	//fill slice with true amount of records which equals to quantity of black holes
	for n := 0; n < g.BlackHoleNum; n++ {
		blackHoleLocations[n] = true
	}

	//shuffle the blackHoleLocations slice to randomly distribute black holes across it
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(blackHoleLocations), func(i, j int) {
		blackHoleLocations[i], blackHoleLocations[j] = blackHoleLocations[j], blackHoleLocations[i]
	})

	return blackHoleLocations
}
