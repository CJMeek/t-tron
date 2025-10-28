package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestGameInitialization(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatalf("failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	player1 := Player{X: 2, Y: 1, Colour: tcell.ColorPurple}
	player2 := Player{X: 8, Y: 1, Colour: tcell.ColorMediumTurquoise}

	game := Game{
		Screen:  screen,
		Player1: player1,
		Player2: player2,
	}

	if game.Player1.X != 2 || game.Player1.Y != 1 {
		t.Errorf("Player1 position incorrect, got (%d, %d)", game.Player1.X, game.Player1.Y)
	}

	if game.Player2.X != 8 || game.Player2.Y != 1 {
		t.Errorf("Player2 position incorrect, got (%d, %d)", game.Player2.X, game.Player2.Y)
	}
}
