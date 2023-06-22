package game

func (g *Game) GetGameStatus() string {
	var gameStatus string
	if g.Status == StatusWon {
		gameStatus = "won"
	} else if g.Status == StatusLost {
		gameStatus = "lost"
	}

	return gameStatus
}

func (g *Game) refreshStatus(lost bool, totalRevealed, totalCells int) {
	if lost {
		g.Status = StatusLost
	} else if totalRevealed+g.BlackHoleNum == totalCells {
		g.Status = StatusWon
	} else {
		g.Status = StatusInProgress
	}
}
