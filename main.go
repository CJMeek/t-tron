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

	player1 := Player{
		X:      2,
		Y:      1,
		Colour: tcell.ColorPurple,
	}

	player2 := Player{
		X:      8,
		Y:      1,
		Colour: tcell.ColorMediumTurquoise,
	}

	game := Game{
		Screen:  s,
		Player1: player1,
		Player2: player2,
	}

	quit := make(chan struct{})
	go game.Run(quit)

	for {

		width, height := s.Size()

		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC || event.Rune() == 'q' {
				close(quit)
				return
			}

			if event.Key() == tcell.KeyUp {
				game.Player1.MoveUp()
			} else if event.Key() == tcell.KeyDown {
				game.Player1.MoveDown(height)
			} else if event.Key() == tcell.KeyLeft {
				game.Player1.MoveLeft()
			} else if event.Key() == tcell.KeyRight {
				game.Player1.MoveRight(width)
			}

			if event.Rune() == 'w' {
				game.Player2.MoveUp()
			} else if event.Rune() == 'a' {
				game.Player2.MoveLeft()
			} else if event.Rune() == 's' {
				game.Player2.MoveDown(height)
			} else if event.Rune() == 'd' {
				game.Player2.MoveRight(width)
			}
		}
	}

}
