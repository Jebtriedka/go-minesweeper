package main

import (
	"log"

	game "github.com/Jebtriedka/minesweeper/Game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenWidth)
	ebiten.SetWindowTitle("Minesweeper - GO")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
