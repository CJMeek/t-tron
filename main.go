package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defer s.Fini()

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	s.SetStyle(defStyle)

	width, height := s.Size()

	player1 := Player{
		X:         1,
		Y:         1,
		Direction: DirectionRight,
		Colour:    tcell.ColorPurple,
	}

	player2 := Player{
		X:         width - 1 - PlayerWidth,
		Y:         height - 1,
		Direction: DirectionLeft,
		Colour:    tcell.ColorMediumTurquoise,
	}

	game := NewGame(s, player1, player2)

	quit := make(chan struct{})
	go game.Run(quit)

	for {

		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC || event.Rune() == 'q' {
				close(quit)
				return
			}

			// Player 1 controls (Arrow keys) - only change direction
			if event.Key() == tcell.KeyUp {
				game.SetPlayer1Direction(DirectionUp)
			} else if event.Key() == tcell.KeyDown {
				game.SetPlayer1Direction(DirectionDown)
			} else if event.Key() == tcell.KeyLeft {
				game.SetPlayer1Direction(DirectionLeft)
			} else if event.Key() == tcell.KeyRight {
				game.SetPlayer1Direction(DirectionRight)
			}

			// Player 2 controls (WASD) - only change direction
			if event.Rune() == 'w' || event.Rune() == 'W' {
				game.SetPlayer2Direction(DirectionUp)
			} else if event.Rune() == 'a' || event.Rune() == 'A' {
				game.SetPlayer2Direction(DirectionLeft)
			} else if event.Rune() == 's' || event.Rune() == 'S' {
				game.SetPlayer2Direction(DirectionDown)
			} else if event.Rune() == 'd' || event.Rune() == 'D' {
				game.SetPlayer2Direction(DirectionRight)
			}
		}
	}

}
