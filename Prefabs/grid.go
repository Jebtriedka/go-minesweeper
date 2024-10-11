package prefabs

import (
	concurrencyPool "github.com/Jebtriedka/minesweeper/ConcurrencyPool"
)

type Grid struct {
	NumUnopenedCells int

	rows  int
	cols  int
	cells []Cell
}

func (grid *Grid) GetCell(i int, j int) *Cell {
	index := grid.getIndex(i, j)
	return &grid.cells[index]
}

func (grid *Grid) GetCells() *[]Cell {
	return &grid.cells
}

func (grid *Grid) GetRow() int {
	return grid.rows
}

func (grid *Grid) GetCol() int {
	return grid.cols
}

func (grid *Grid) SetGridData(numOfCols int, numOfRows int) {
	grid.rows = numOfRows
	grid.cols = numOfCols
	grid.NumUnopenedCells = numOfRows * numOfCols

	grid.cells = make([]Cell, numOfRows*numOfCols)
}

func (grid *Grid) getIndex(i int, j int) (index int) {

	if i < 0 || i >= grid.rows || j < 0 || j >= grid.cols {

	}

	return i*grid.cols + j
}

func (grid *Grid) BfsReveal(startRowIndex, startColIndex int) {

	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	cell := grid.GetCell(startRowIndex, startColIndex)

	if !cell.Revealed() {
		cell.Reveale()
		grid.NumUnopenedCells -= 1
	}

	for i := 0; i < 8; i++ {
		ni := startRowIndex + dx[i]
		nj := startColIndex + dy[i]

		if ni >= 0 && ni < grid.GetRow() && nj >= 0 && nj < grid.GetCol() {

			if grid.GetCell(ni, nj).Revealed() || cell.Value != "0" {
				continue
			}

			grid.BfsReveal(ni, nj)
		}
	}
}

func (grid *Grid) RevealeAllMines() {

	sliceSize := len(grid.cells) / 4
	var tasks []concurrencyPool.Task

	for i := 0; i < 4; i++ {
		start_index := i * sliceSize
		end_index := start_index + sliceSize

		tasks = append(tasks, &RevealeAllMinesTask{
			Cells:      grid.cells,
			StartIndex: start_index,
			EndIndex:   end_index})
	}

	wp := concurrencyPool.WorkerPool{
		Tasks:       tasks,
		Concurrency: 3,
	}

	wp.Run()
}
