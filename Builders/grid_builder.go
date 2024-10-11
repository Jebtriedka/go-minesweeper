package builders

import (
	"image/color"

	concurrencyPool "github.com/Jebtriedka/minesweeper/ConcurrencyPool"
	prefabs "github.com/Jebtriedka/minesweeper/Prefabs"
)

func GridBuilder(rows int, cols int) prefabs.Grid {

	if rows <= 0 || cols <= 0 {

	}

	grid := prefabs.Grid{}
	grid.SetGridData(rows, cols)
	cells := grid.GetCells()

	sliceSize := len(*cells) / 4
	var tasks []concurrencyPool.Task

	for i := 0; i < 4; i++ {
		start_index := i * sliceSize
		end_index := start_index + sliceSize

		tasks = append(tasks, &BuildGridTask{
			Cells:      *cells,
			StartIndex: start_index,
			EndIndex:   end_index})
	}

	wp := concurrencyPool.WorkerPool{
		Tasks:       tasks,
		Concurrency: 3,
	}

	wp.Run()

	return grid
}

type BuildGridTask struct {
	Cells      []prefabs.Cell
	StartIndex int
	EndIndex   int
}

func (t *BuildGridTask) Process() {
	for i := t.StartIndex; i < t.EndIndex; i++ {
		t.Cells[i].Value = ""
		t.Cells[i].BackGroundColor = color.White
	}
}
