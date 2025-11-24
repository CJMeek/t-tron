package main

import (
	"testing"
	"time"

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

	game := NewGame(screen, player1, player2)

	if game.Player1.X != 2 || game.Player1.Y != 1 {
		t.Errorf("Player1 position incorrect, got (%d, %d)", game.Player1.X, game.Player1.Y)
	}

	if game.Player2.X != 8 || game.Player2.Y != 1 {
		t.Errorf("Player2 position incorrect, got (%d, %d)", game.Player2.X, game.Player2.Y)
	}
}

func TestGameRunAutoMovesPlayersByDirection(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatalf("failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	p1 := Player{X: 4, Y: 2, Direction: DirectionRight, Colour: tcell.ColorPurple}
	p2 := Player{X: 10, Y: 2, Direction: DirectionLeft, Colour: tcell.ColorMediumTurquoise}
	game := NewGame(screen, p1, p2)

	startX1 := game.Player1.X
	startX2 := game.Player2.X

	quit := make(chan struct{})
	go game.Run(quit)

	time.Sleep(35 * time.Millisecond) // allow a couple of ticks

	close(quit)
	time.Sleep(5 * time.Millisecond) // let goroutine exit

	if game.Player1.X <= startX1 {
		t.Errorf("expected Player1 to move right from %d to > %d, got %d", startX1, startX1, game.Player1.X)
	}
	if game.Player2.X >= startX2 {
		t.Errorf("expected Player2 to move left from %d to < %d, got %d", startX2, startX2, game.Player2.X)
	}
}

func TestGameInitializationWithDirection(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatalf("failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	player1 := Player{X: 2, Y: 1, Direction: DirectionRight, Colour: tcell.ColorPurple}
	player2 := Player{X: 8, Y: 1, Direction: DirectionLeft, Colour: tcell.ColorMediumTurquoise}

	game := NewGame(screen, player1, player2)

	if game.Player1.Direction != DirectionRight {
		t.Errorf("expected Player1 direction DirectionRight, got %v", game.Player1.Direction)
	}
	if game.Player2.Direction != DirectionLeft {
		t.Errorf("expected Player2 direction DirectionLeft, got %v", game.Player2.Direction)
	}
}

func TestGameSetPlayerDirectionThreadSafe(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatalf("failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	player1 := Player{X: 2, Y: 1, Direction: DirectionRight, Colour: tcell.ColorPurple}
	player2 := Player{X: 8, Y: 1, Direction: DirectionLeft, Colour: tcell.ColorMediumTurquoise}
	game := NewGame(screen, player1, player2)

	// Test that SetPlayer1Direction works
	game.SetPlayer1Direction(DirectionUp)
	if game.Player1.Direction != DirectionUp {
		t.Errorf("expected Player1 direction to be DirectionUp, got %v", game.Player1.Direction)
	}

	// Test that SetPlayer2Direction works
	game.SetPlayer2Direction(DirectionDown)
	if game.Player2.Direction != DirectionDown {
		t.Errorf("expected Player2 direction to be DirectionDown, got %v", game.Player2.Direction)
	}
}
