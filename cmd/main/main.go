package main

import (
	mouse "github.com/HMasataka/mouse_gesture"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	mouseGesture mouse.MouseGesture
}

func (g *Game) Update() error {
	position := mouse.NewPosition(ebiten.CursorPosition())

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.mouseGesture.Record(position)
	}

	if g.mouseGesture.IsEnable() && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.mouseGesture.Finish()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	var mouseGesture mouse.MouseGesture

	if err := ebiten.RunGame(&Game{
		mouseGesture: mouseGesture,
	}); err != nil {
		panic(err)
	}
}
