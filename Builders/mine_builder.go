package builders

import (
	"math/rand"
	"time"

	prefabs "github.com/Jebtriedka/minesweeper/Prefabs"
)

func MineBuilder(grid *prefabs.Grid, mineCount int) *prefabs.Grid {

	totalCells := grid.GetRow() * grid.GetCol()
	if mineCount > totalCells {

	}

	for mineCount > 0 {
		mineCount--

		for {

			row, col := pickRandomCell(grid.GetRow(), grid.GetCol())
			cell := grid.GetCell(row, col)

			if cell.IsMine() {
				continue
			}

			cell.PlaceMine()
			updateNumberOfNeighbourMines(grid, row, col)

			break
		}
	}

	return grid
}

func updateNumberOfNeighbourMines(grid *prefabs.Grid, rowMineIndex int, colMineIndex int) {

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		ni := rowMineIndex + dx[i]
		nj := colMineIndex + dy[i]

		if ni >= 0 && ni < grid.GetRow() && nj >= 0 && nj < grid.GetCol() {
			cell := grid.GetCell(ni, nj)
			cell.SetNumberOfNeighbourMines(cell.GetNumberOfNeighbourMines() + 1)

		}
	}
}

func pickRandomCell(MaxRows int, MaxCols int) (int, int) {

	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(MaxRows)
	j := rand.Intn(MaxCols)

	return i, j
}
