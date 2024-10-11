package prefabs

type ICell interface {
	Get_number_of_neighbour_mines() int
	IsMine() bool
	Place_mine()
	Set_number_of_neighbour_mines(numberOfNeighbourMines int)
}
