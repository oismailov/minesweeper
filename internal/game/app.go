package game

import "minesweeper/internal/cell"

func StartNewGame(height, width, blackHolesNum int) *Game {
	grid := make([][]*cell.Cell, height)

	// fill in the Grid
	for i := 0; i < height; i++ {
		grid[i] = make([]*cell.Cell, width)

		for j := 0; j < width; j++ {
			grid[i][j] = cell.NewCell(cell.TypeBlank, i, j)
		}
	}

	//create new game and fill the grid with randomly located black holes
	game := NewGame(grid, StatusInProgress, height, width, blackHolesNum)
	game.plantBlackHoles()

	return game
}
