package game

import (
	"os"

	builders "github.com/Jebtriedka/minesweeper/Builders"
	prefabs "github.com/Jebtriedka/minesweeper/Prefabs"

	"github.com/hajimehoshi/ebiten/v2"
)

var gird prefabs.Grid

type Game struct {
	state       int
	winMessage  string
	lostMessage string
}

func NewGame() *Game {

	gird = builders.GridBuilder(gridSize, gridSize)
	builders.MineBuilder(&gird, mineCount)

	return &Game{
		state:       Playing,
		winMessage:  winMessage,
		lostMessage: lostMessage,
	}
}

func init() {
	NewGame()
}

func (g *Game) Update() error {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.state == Playing {
		x, y := ebiten.CursorPosition()

		i := y / cellSize
		j := x / cellSize

		if i >= 0 && i < gridSize && j >= 0 && j < gridSize {
			cell := gird.GetCell(i, j)
			if !cell.Revealed() {

				if cell.IsMine() {
					gird.RevealeAllMines()
					g.state = GameOver
				} else {
					gird.BfsReveal(i, j)
				}
			}
		}

		// fmt.Println(gird.NumUnopenedCells)
	}

	if gird.NumUnopenedCells == mineCount {
		g.state = Win
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		// fmt.Println("Exit the game")
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {

		// fmt.Println("Restart the game")
		g.state = Playing
		NewGame()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.PlayingScreen(screen)
	switch g.state {
	case Win:
		g.RenderPopUp(screen, g.winMessage)
	case GameOver:
		g.RenderPopUp(screen, g.lostMessage)
	}
}
