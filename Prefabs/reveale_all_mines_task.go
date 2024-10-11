package prefabs

type RevealeAllMinesTask struct {
	Cells      []Cell
	StartIndex int
	EndIndex   int
}

func (t *RevealeAllMinesTask) Process() {
	for i := t.StartIndex; i < t.EndIndex; i++ {
		if t.Cells[i].IsMine() {
			t.Cells[i].Reveale()
		}
	}
}
