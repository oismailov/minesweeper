package game

import c "minesweeper/internal/cell"

func (g *Game) refreshState() {
	var (
		totalRevealed int
		lost          bool
	)

	totalCells := g.Height * g.Width
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			cell := g.Grid[i][j]
			if cell.IsOpen {
				totalRevealed++
			}

			if cell.IsOpen && cell.Type == c.TypeBlackHole {
				lost = true
			}
		}
	}

	g.refreshStatus(lost, totalRevealed, totalCells)
}
