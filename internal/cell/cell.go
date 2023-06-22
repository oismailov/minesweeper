package cell

type Type int

const (
	TypeBlackHole Type = iota
	TypeBlank
)

type Cell struct {
	X                  int
	Y                  int
	IsOpen             bool
	AdjacentBlackHoles int
	Type               Type
}

func NewCell(cellType Type, x, y int) *Cell {
	return &Cell{
		X:                  x,
		Y:                  y,
		IsOpen:             false,
		AdjacentBlackHoles: 0,
		Type:               cellType,
	}
}

func (c *Cell) Open() {
	c.IsOpen = true
}

func (c *Cell) GetAdjacentCells(grid [][]*Cell) []*Cell {
	var adjacentCells []*Cell

	//each cell will have 8 adjacent cells in two-dimensional space
	//{x,y} is location of adjacent cell where {0,0} is current cell location
	adjacentCellsLocations := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, adjacentCellsLocation := range adjacentCellsLocations {
		dx, dy := adjacentCellsLocation[0], adjacentCellsLocation[1]
		x, y := c.X+dx, c.Y+dy

		//check if adjacent cell do not cross the Grid (is not outside the Grid)
		if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) {
			adjacentCells = append(adjacentCells, grid[x][y])
		}
	}

	return adjacentCells
}
