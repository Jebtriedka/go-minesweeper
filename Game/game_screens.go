package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) PlayingScreen(screen *ebiten.Image) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {

			ebitenutil.DrawRect(screen, float64(j*cellSize), float64(i*cellSize), cellSize, cellSize, gird.GetCell(i, j).BackGroundColor)

			ebitenutil.DrawRect(screen, float64(j*cellSize), float64(i*cellSize), cellSize, 1, color.Black)
			ebitenutil.DrawRect(screen, float64(j*cellSize), float64(i*cellSize+cellSize-1), cellSize, 1, color.Black)
			ebitenutil.DrawRect(screen, float64(j*cellSize), float64(i*cellSize), 1, cellSize, color.Black)
			ebitenutil.DrawRect(screen, float64(j*cellSize+cellSize-1), float64(i*cellSize), 1, cellSize, color.Black)

			valueStr := gird.GetCell(i, j).Value

			textX := float64(j*cellSize + (cellSize/2 - 8))
			textY := float64(i*cellSize + (cellSize/2 + 8))

			ebitenutil.DebugPrintAt(screen, valueStr, int(textX), int(textY))

		}
	}
}

func (g *Game) RenderPopUp(screen *ebiten.Image, message string) {
	popupX := (ScreenWidth - popupWidth) / 2
	popupY := (ScreenHeight - popupHeight) / 2

	ebitenutil.DrawRect(screen, float64(popupX-borderWidth), float64(popupY-borderWidth), popupWidth+2*borderWidth, popupHeight+2*borderWidth, color.RGBA{255, 215, 0, 255}) // Gold border
	ebitenutil.DrawRect(screen, float64(popupX), float64(popupY), popupWidth, popupHeight, color.RGBA{0, 0, 0, 200})                                                         // Black background with transparency

	textX := popupX + (popupWidth / 2) - (len(message) * 3)
	textY := popupY + (popupHeight / 2) - 10
	ebitenutil.DebugPrintAt(screen, message, textX, textY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
