package game

const (
	winMessage  = "You Won! Press R to restart or Q to quit. "
	lostMessage = "You Lost! Press R to restart or Q to quit."
)

const (
	mineCount = 12
	gridSize  = 10
)

const (
	Playing = iota
	Win
	GameOver
)

const (
	cellSize     = 64
	ScreenWidth  = gridSize * cellSize
	ScreenHeight = gridSize * cellSize
)

const (
	screenWidth  = 640
	screenHeight = 480
	popupWidth   = 300
	popupHeight  = 150
	borderWidth  = 4
)
