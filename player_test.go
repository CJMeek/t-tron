package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestPlayerInitialization(t *testing.T) {
	player := Player{X: 5, Y: 5, Direction: DirectionRight, Colour: tcell.ColorRed}
	if player.X != 5 {
		t.Errorf("expected X to be 5, got %d", player.X)
	}
	if player.Y != 5 {
		t.Errorf("expected Y to be 5, got %d", player.Y)
	}
	if player.Direction != DirectionRight {
		t.Errorf("expected Direction to be DirectionRight, got %v", player.Direction)
	}
	if player.Colour != tcell.ColorRed {
		t.Errorf("expected Colour to be Red, got %v", player.Colour)
	}
}

func TestPlayerSetDirectionAllowsValidTurns(t *testing.T) {
	player := Player{X: 5, Y: 5, Direction: DirectionRight, Colour: tcell.ColorRed}

	// Moving up from right should work
	player.SetDirection(DirectionUp)
	if player.Direction != DirectionUp {
		t.Errorf("expected Direction to be DirectionUp, got %v", player.Direction)
	}

	// Moving left from up should work
	player.SetDirection(DirectionLeft)
	if player.Direction != DirectionLeft {
		t.Errorf("expected Direction to be DirectionLeft, got %v", player.Direction)
	}

	// Moving down from left should work
	player.SetDirection(DirectionDown)
	if player.Direction != DirectionDown {
		t.Errorf("expected Direction to be DirectionDown, got %v", player.Direction)
	}
}

func TestPlayerSetDirectionPrevents180DegreeTurns(t *testing.T) {
	player := Player{X: 5, Y: 5, Direction: DirectionRight, Colour: tcell.ColorRed}

	// Try to go left while moving right (should be blocked)
	player.SetDirection(DirectionLeft)
	if player.Direction != DirectionRight {
		t.Errorf("expected Direction to remain DirectionRight, got %v", player.Direction)
	}

	// Change to up
	player.SetDirection(DirectionUp)

	// Try to go down while moving up (should be blocked)
	player.SetDirection(DirectionDown)
	if player.Direction != DirectionUp {
		t.Errorf("expected Direction to remain DirectionUp, got %v", player.Direction)
	}
}

func TestPlayerSetDirectionAllDirections(t *testing.T) {
	tests := []struct {
		name         string
		currentDir   Direction
		newDir       Direction
		shouldChange bool
		expectedDir  Direction
	}{
		{"Up to Down blocked", DirectionUp, DirectionDown, false, DirectionUp},
		{"Up to Left allowed", DirectionUp, DirectionLeft, true, DirectionLeft},
		{"Up to Right allowed", DirectionUp, DirectionRight, true, DirectionRight},
		{"Down to Up blocked", DirectionDown, DirectionUp, false, DirectionDown},
		{"Down to Left allowed", DirectionDown, DirectionLeft, true, DirectionLeft},
		{"Down to Right allowed", DirectionDown, DirectionRight, true, DirectionRight},
		{"Left to Right blocked", DirectionLeft, DirectionRight, false, DirectionLeft},
		{"Left to Up allowed", DirectionLeft, DirectionUp, true, DirectionUp},
		{"Left to Down allowed", DirectionLeft, DirectionDown, true, DirectionDown},
		{"Right to Left blocked", DirectionRight, DirectionLeft, false, DirectionRight},
		{"Right to Up allowed", DirectionRight, DirectionUp, true, DirectionUp},
		{"Right to Down allowed", DirectionRight, DirectionDown, true, DirectionDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := Player{X: 5, Y: 5, Direction: tt.currentDir, Colour: tcell.ColorRed}
			player.SetDirection(tt.newDir)
			if player.Direction != tt.expectedDir {
				t.Errorf("expected Direction to be %v, got %v", tt.expectedDir, player.Direction)
			}
		})
	}
}
