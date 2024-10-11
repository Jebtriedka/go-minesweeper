package prefabs

import (
	"image/color"
	"strconv"
)

type Cell struct {
	Value           string
	BackGroundColor color.Color

	cellType               CellType
	numberOfNeighbourMines int
}

func (c *Cell) GetNumberOfNeighbourMines() int {
	return c.numberOfNeighbourMines
}

func (c *Cell) SetNumberOfNeighbourMines(numberOfNeighbourMines int) {
	c.numberOfNeighbourMines = numberOfNeighbourMines
}

func (c *Cell) PlaceMine() {

	c.cellType = Mine

}

func (c *Cell) IsMine() bool {
	return c.cellType == Mine
}

func (c *Cell) Revealed() bool {
	return c.BackGroundColor == color.Black
}

func (c *Cell) Reveale() {
	c.BackGroundColor = color.Black
	if c.cellType == Empty {
		c.Value = strconv.Itoa(c.numberOfNeighbourMines)
	} else {
		c.Value = " 😀 "
		c.BackGroundColor = color.RGBA{255, 0, 0, 0}
	}
}
