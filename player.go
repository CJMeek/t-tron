package main

import (
	"github.com/gdamore/tcell/v2"
)

// Direction represents the direction a player is moving
type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

// PlayerWidth is the width of a player sprite in cells
const PlayerWidth = 2

type Player struct {
	X         int
	Y         int
	Direction Direction
	Colour    tcell.Color
}

// SetDirection changes the player's direction, preventing 180-degree turns
func (p *Player) SetDirection(dir Direction) {
	// Prevent 180-degree turns
	if (p.Direction == DirectionUp && dir == DirectionDown) ||
		(p.Direction == DirectionDown && dir == DirectionUp) ||
		(p.Direction == DirectionLeft && dir == DirectionRight) ||
		(p.Direction == DirectionRight && dir == DirectionLeft) {
		return // ignore the input
	}
	p.Direction = dir
}
