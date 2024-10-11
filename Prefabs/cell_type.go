package prefabs

type CellType int

const (
	Empty CellType = iota
	Mine
)

func (s CellType) String() string {
	return [...]string{"Empty", "Mine"}[s]
}
