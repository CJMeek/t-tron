package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen  tcell.Screen
	Player1 Player
	Player2 Player
}

func (g *Game) Run(quit <-chan struct{}) {

	//painted map to show where trails are as to avoid collisions

	// make players move constantly
	// change keys purely to change direction

	//if direction = up, or down then y-- or y++
	//if direciton = left or right then x-- or x++
	//may be an edge case where player suddenly moves diagonally
	//edge case where players move into the same space at the same time or die at the same time -> need to add flag for both players crashing every tick

	//create a draw function that draws the players and their trails
	//fills in gaps on lower framerates -> seeing issue with 30fps (33ms ticker). Could be a problem for online multiplayer

	playerStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	trailStyle1 := tcell.StyleDefault.Background(g.Player1.Colour).Foreground(g.Player1.Colour)
	trailStyle2 := tcell.StyleDefault.Background(g.Player2.Colour).Foreground(g.Player2.Colour)
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	lastX1, lastY1 := g.Player1.X, g.Player1.Y
	lastX2, lastY2 := g.Player2.X, g.Player2.Y

	for {

		select {
		case <-quit:
			return
		case <-ticker.C:

			if lastX1 != g.Player1.X || lastY1 != g.Player1.Y {
				g.Screen.SetContent(lastX1, lastY1, ' ', nil, trailStyle1)
				g.Screen.SetContent(lastX1+1, lastY1, ' ', nil, trailStyle1)
				lastX1, lastY1 = g.Player1.X, g.Player1.Y
			}
			g.Screen.SetContent(g.Player1.X, g.Player1.Y, ' ', nil, playerStyle)
			g.Screen.SetContent(g.Player1.X+1, g.Player1.Y, ' ', nil, playerStyle)

			if lastX2 != g.Player2.X || lastY2 != g.Player2.Y {
				g.Screen.SetContent(lastX2, lastY2, ' ', nil, trailStyle2)
				g.Screen.SetContent(lastX2+1, lastY2, ' ', nil, trailStyle2)
				lastX2, lastY2 = g.Player2.X, g.Player2.Y
			}
			g.Screen.SetContent(g.Player2.X, g.Player2.Y, ' ', nil, playerStyle)
			g.Screen.SetContent(g.Player2.X+1, g.Player2.Y, ' ', nil, playerStyle)

			g.Screen.Show()
		}

	}

}

// func (g *Game) drawSprite(x int, y int, )
