package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

func (g *Game) Update() error {
	println(ebiten.CursorPosition())
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
