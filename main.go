package main

import (
	"github.com/HMasataka/mouse_gesture/mouse"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	mouseGesture mouse.MouseGesture
}

func (g *Game) Update() error {
	position := mouse.NewPosition(ebiten.CursorPosition())

	println(position.X, position.Y)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		println("Left mouse button pressed")
	}

	println(g.mouseGesture.GetDirection(position))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	var mouseGesture mouse.MouseGesture
	mouseGesture.Enable()

	if err := ebiten.RunGame(&Game{
		mouseGesture: mouseGesture,
	}); err != nil {
		panic(err)
	}
}
