package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
	figureWidth  = 100
	figureHeight = 300
)

var (
	figureX = (screenWidth - figureWidth) / 2
	figureY = (screenHeight - figureHeight) / 2
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		figureX = x - figureWidth/2
		figureY = y - figureHeight/2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	drawFigure(screen, figureX, figureY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func drawFigure(screen *ebiten.Image, x, y int) {
	figureColor := color.RGBA{255, 255, 0, 255}
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(figureWidth), float64(figureHeight), figureColor)
	ebitenutil.DrawRect(screen, float64(x), float64(y+figureHeight/2-figureWidth/2), float64(figureWidth), float64(figureHeight), figureColor)
	ebitenutil.DrawRect(screen, float64(x+figureWidth/2-figureHeight/2), float64(y), float64(figureHeight), float64(figureWidth), figureColor)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Painter")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
