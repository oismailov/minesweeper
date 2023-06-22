package game

import (
	"errors"
	c "minesweeper/internal/cell"
)

func (g *Game) RevealCell(x, y int) ([]*c.Cell, error) {
	if g.Status != StatusInProgress {
		return nil, errors.New("game is over")
	}

	if err := g.validateCoords(x, y); err != nil {
		return []*c.Cell{}, err
	}

	cell := g.Grid[y][x]
	//there is no need to reveal already revealed cell
	if cell.IsOpen {
		return []*c.Cell{cell}, nil
	}

	// if cell type is a back hole or number of adjacent black holes is more than 0
	//then there is no need to traverse adjacent cells
	if cell.Type == c.TypeBlackHole || cell.AdjacentBlackHoles > 0 {
		cell.Open()
		g.refreshState()

		return []*c.Cell{cell}, nil
	}

	//if cell is not a black hole and number of adjacent black holes is equal to zero
	//then adjacent cells can be traversed
	revealedCells := g.traverseAdjacentCells(cell)
	g.refreshState()

	return revealedCells, nil
}

func (g *Game) validateCoords(x, y int) error {
	if x < 0 || x >= g.Height || y < 0 || y >= g.Width {
		return errors.New("invalid cell")
	}

	return nil
}

// traverse adjacent cells that can be revealed
func (g *Game) traverseAdjacentCells(cell *c.Cell) []*c.Cell {
	var revealedCells []*c.Cell
	//amount of cells that is going to be discovered
	discovered := make(map[*c.Cell]bool)

	//queue that keeps the number of all cells
	//channel structure works well here, because it is easy to manage read and write operations
	queue := make(chan *c.Cell, g.Height*g.Width)

	queue <- cell
	discovered[cell] = true
	for len(queue) > 0 {
		curCell := <-queue
		curCell.Open()
		revealedCells = append(revealedCells, curCell)

		if curCell.AdjacentBlackHoles == 0 {
			for _, adjacentCell := range curCell.GetAdjacentCells(g.Grid) {
				if !adjacentCell.IsOpen && !discovered[adjacentCell] {
					discovered[adjacentCell] = true
					queue <- adjacentCell
				}
			}
		}
	}

	return revealedCells
}
