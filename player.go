package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	X         int
	Y         int
	Direction string
	Colour    tcell.Color
}

func (p *Player) MoveUp() {
	if p.Direction != "up" {
		if p.Y > 0 {
			p.Y -= 1
			p.Direction = "up"
		}
	}
}

func (p *Player) MoveDown(maxHeight int) {
	if p.Direction != "down" {
		if p.Y < maxHeight-1 {
			p.Y += 1
			p.Direction = "down"
		}
	}

}

func (p *Player) MoveLeft() {
	if p.Direction != "left" {
		if p.X > 0 {
			p.X -= 2
			p.Direction = "left"
		}
	}
}

func (p *Player) MoveRight(maxWidth int) {
	if p.Direction != "right" {
		if p.X < maxWidth-2 {
			p.X += 2
			p.Direction = "right"
		}
	}

}

func (p *Player) Display() string {
	return strings.Repeat(" ", 2)
}
