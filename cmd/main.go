package main

import (
	"bufio"
	"fmt"
	c "minesweeper/internal/cell"
	g "minesweeper/internal/game"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter grid height: ")
	input, _, _ := reader.ReadLine()
	height, _ := strconv.Atoi(string(input))

	fmt.Print("Enter grid width: ")
	input, _, _ = reader.ReadLine()
	width, _ := strconv.Atoi(string(input))

	fmt.Print("Enter amount of black holes in grid: ")
	input, _, _ = reader.ReadLine()
	blackHolesNum, _ := strconv.Atoi(string(input))

	game := g.StartNewGame(height, width, blackHolesNum)

	printGrid(game, height, width)

	for {
		//get coordinates from CLI
		x, y := getCoordinates(reader)

		_, err := game.RevealCell(x, y)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("\n")
		printGrid(game, height, width)

		//play until game status is in progress
		if game.Status != g.StatusInProgress {
			break
		}
	}

	//print human-readable game status
	gameStatus := game.GetGameStatus()
	fmt.Println("You ", gameStatus, " the game!")

}

func getCoordinates(reader *bufio.Reader) (int, int) {
	fmt.Print("Enter coordinates separated by comma: ")
	//read coordinates from CLI
	input, _, _ := reader.ReadLine()
	coordinates := strings.Split(string(input), ",")

	//convert input to integer
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])

	return x, y
}

func printGrid(game *g.Game, height, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cell := game.Grid[i][j]
			//green (revealed) cell
			if cell.IsOpen && cell.Type == c.TypeBlank && cell.AdjacentBlackHoles == 0 {
				fmt.Print("\U0001F7E9", "\t")
			}

			//show number of adjacent black holes
			if cell.IsOpen && cell.Type == c.TypeBlank && cell.AdjacentBlackHoles > 0 {
				fmt.Print(cell.AdjacentBlackHoles, "\t")
			}

			//black (mine) cell
			if cell.IsOpen && cell.Type == c.TypeBlackHole {
				fmt.Print("⬛", "\t")
			}

			//cell that has not been revealed yed
			if !cell.IsOpen {
				fmt.Print("⬜", "\t")
			}
		}

		fmt.Println("\n")
	}
}
