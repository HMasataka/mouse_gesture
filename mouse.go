package mousegesture

import (
	"fmt"
	"reflect"
)

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
	gesture  Gesture
}

func (m *MouseGesture) Enable(startPosition Position) {
	m.old = startPosition
	m.gesture = Gesture{}
	m.isEnable = true
}

func (m *MouseGesture) Disable() {
	m.isEnable = false
}

func (m *MouseGesture) IsEnable() bool {
	return m.isEnable
}

func (m *MouseGesture) getDirection(pos Position) Direction {
	if !m.isEnable {
		return None
	}

	direction := None

	isX := Abs(m.old.X-pos.X) > Abs(m.old.Y-pos.Y)

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

func (m *MouseGesture) Record(pos Position) {
	if !m.isEnable {
		m.Enable(pos)
	}

	direction := m.getDirection(pos)

	if direction == None {
		return
	}

	if len(m.gesture) > 0 && m.gesture[len(m.gesture)-1] == direction {
		return
	}

	m.gesture = append(m.gesture, direction)
}

func (m *MouseGesture) Finish() Gesture {
	m.Disable()

	for pair := AllGestures.Oldest(); pair != nil; pair = pair.Next() {
		if reflect.DeepEqual(pair.Value, m.gesture) {
			fmt.Printf("%s => %v\n", pair.Key, pair.Value)
		}
	}

	return []Direction{}
}
