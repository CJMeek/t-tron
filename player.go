package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	X      int
	Y      int
	Colour tcell.Color
}

func (p *Player) MoveUp() {
	if p.Y > 0 {
		p.Y -= 1
	}

}

func (p *Player) MoveDown(maxHeight int) {

	if p.Y < maxHeight-1 {
		p.Y += 1
	}

}

func (p *Player) MoveLeft() {
	if p.X > 0 {
		p.X -= 2
	}
}

func (p *Player) MoveRight(maxWidth int) {
	if p.X < maxWidth-2 {
		p.X += 2
	}

}

func (p *Player) Display() string {
	return strings.Repeat(" ", 2)
}
