package core

import (
	"sync"
)

type Lawn struct {
	Mtx           sync.Mutex
	Width, Height int

	Matrix map[Position]int8
}

func InitLawn(w, h int) *Lawn {
	m := make(map[Position]int8)

	return &Lawn{
		Width:  w,
		Height: h,
		Matrix: m,
	}
}

func (l *Lawn) IsValidPosition(pos Position) bool {
	return pos.X >= 0 && pos.X <= l.Width && pos.Y >= 0 && pos.Y <= l.Height
}
