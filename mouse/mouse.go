package mouse

import "github.com/HMasataka/mouse_gesture/math"

const threshold = 5

func NewPosition(x, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

type Position struct {
	X int
	Y int
}

type Direction int

const (
	None Direction = iota
	Left
	Right
	Up
	Down
)

type MouseGesture struct {
	isEnable bool
	old      Position
}

func (m *MouseGesture) Enable() {
	m.isEnable = true
}

func (m *MouseGesture) Disable() {
	m.isEnable = false
}

func (m *MouseGesture) IsEnable() bool {
	return m.isEnable
}

func (m *MouseGesture) GetDirection(pos Position) Direction {
	if !m.isEnable {
		return None
	}

	direction := None

	isX := math.Abs(m.old.X-pos.X) > math.Abs(m.old.Y-pos.Y)

	if isX {
		if m.old.X-pos.X > threshold {
			direction = Left
		} else if pos.X-m.old.X > threshold {
			direction = Right
		}
	} else {
		if m.old.Y-pos.Y > threshold {
			direction = Up
		} else if pos.Y-m.old.Y > threshold {
			direction = Down
		}
	}

	m.old = pos

	return direction
}
