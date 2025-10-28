package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestPlayerMoveUp(t *testing.T) {
	player := Player{X: 5, Y: 5, Colour: tcell.ColorRed}
	player.MoveUp()
	if player.Y != 4 {
		t.Errorf("expected Y to be 4, got %d", player.Y)
	}
}

func TestPlayerMoveDown(t *testing.T) {
	player := Player{X: 5, Y: 5, Colour: tcell.ColorRed}
	player.MoveDown(10)
	if player.Y != 6 {
		t.Errorf("expected Y to be 6, got %d", player.Y)
	}
}

func TestPlayerMoveLeft(t *testing.T) {
	player := Player{X: 5, Y: 5, Colour: tcell.ColorRed}
	player.MoveLeft()
	if player.X != 3 {
		t.Errorf("expected X to be 3, got %d", player.X)
	}
}

func TestPlayerMoveRight(t *testing.T) {
	player := Player{X: 5, Y: 5, Colour: tcell.ColorRed}
	player.MoveRight(10)
	if player.X != 7 {
		t.Errorf("expected X to be 7, got %d", player.X)
	}
}

func TestPlayerDisplay(t *testing.T) {
	player := Player{X: 5, Y: 5, Colour: tcell.ColorRed}
	display := player.Display()
	if display != "  " {
		t.Errorf("expected display to be '  ', got '%s'", display)
	}
}
