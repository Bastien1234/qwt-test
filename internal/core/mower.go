package core

import (
	"fmt"
	"strings"
)

type Position struct {
	X, Y int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "Unknown"
	}
}

func ParseDirection(s string) (Direction, error) {
	switch strings.ToUpper(s) {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	default:
		return North, fmt.Errorf("invalid direction: %s", s)
	}
}

type Mower struct {
	Id int8
	Position
	Direction
	ActionQueue string
}

func InitMower(id int8, pos Position, direction Direction, actions string) *Mower {
	return &Mower{
		Id:          id,
		Position:    pos,
		Direction:   direction,
		ActionQueue: actions,
	}
}

func (m *Mower) TryMove(lawn *Lawn, from, to Position) bool {
	if !lawn.IsValidPosition(to) {
		fmt.Printf("Error log: invalid position required. X: %d Y: %d\n", to.X, to.Y)
		return false
	}

	lawn.Mtx.Lock()
	defer lawn.Mtx.Unlock()

	_, exist := lawn.Matrix[to]
	if exist {
		// No log in this case
		return false
	}

	lawn.Matrix[to] = m.Id

	delete(lawn.Matrix, from)

	return true
}

func (m *Mower) ExecuteMoving(lawn *Lawn) {
	for _, instruction := range m.ActionQueue {
		switch instruction {
		case 'L':
			m.TurnLeft()
		case 'R':
			m.TurnRight()
		case 'F':
			newPos := m.GetForwardPosition()
			if m.TryMove(lawn, m.Position, newPos) {
				m.Position = newPos
			}
		}
	}
}

func (m *Mower) TurnLeft() {
	m.Direction = Direction((int(m.Direction) + 3) % 4)
}

func (m *Mower) TurnRight() {
	m.Direction = Direction((int(m.Direction) + 1) % 4)
}

func (m *Mower) GetForwardPosition() Position {
	switch m.Direction {
	case North:
		return Position{m.Position.X, m.Position.Y + 1}
	case East:
		return Position{m.Position.X + 1, m.Position.Y}
	case South:
		return Position{m.Position.X, m.Position.Y - 1}
	case West:
		return Position{m.Position.X - 1, m.Position.Y}
	default:
		return m.Position
	}
}

func (m *Mower) PrintFinalPosition() {
	fmt.Printf("[*] Mower id: %d, at X: %d Y: %d and facing: %s [*]\n", m.Id, m.X, m.Y, m.Direction)
}
