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

	paintedMap := make(map[int]map[int]bool)

	playerStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	trailStyle1 := tcell.StyleDefault.Background(g.Player1.Colour).Foreground(g.Player1.Colour)
	trailStyle2 := tcell.StyleDefault.Background(g.Player2.Colour).Foreground(g.Player2.Colour)
	ticker := time.NewTicker(33 * time.Millisecond)
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

				// Mark trail positions as painted
				if paintedMap[lastX1] == nil {
					paintedMap[lastX1] = make(map[int]bool)
				}
				paintedMap[lastX1][lastY1] = true

				if paintedMap[lastX1+1] == nil {
					paintedMap[lastX1+1] = make(map[int]bool)
				}
				paintedMap[lastX1+1][lastY1] = true

				lastX1, lastY1 = g.Player1.X, g.Player1.Y
			}
			g.Screen.SetContent(g.Player1.X, g.Player1.Y, ' ', nil, playerStyle)
			g.Screen.SetContent(g.Player1.X+1, g.Player1.Y, ' ', nil, playerStyle)

			if lastX2 != g.Player2.X || lastY2 != g.Player2.Y {
				g.Screen.SetContent(lastX2, lastY2, ' ', nil, trailStyle2)
				g.Screen.SetContent(lastX2+1, lastY2, ' ', nil, trailStyle2)

				// Mark trail positions as painted
				if paintedMap[lastX2] == nil {
					paintedMap[lastX2] = make(map[int]bool)
				}
				paintedMap[lastX2][lastY2] = true

				if paintedMap[lastX2+1] == nil {
					paintedMap[lastX2+1] = make(map[int]bool)
				}
				paintedMap[lastX2+1][lastY2] = true

				lastX2, lastY2 = g.Player2.X, g.Player2.Y
			}
			g.Screen.SetContent(g.Player2.X, g.Player2.Y, ' ', nil, playerStyle)
			g.Screen.SetContent(g.Player2.X+1, g.Player2.Y, ' ', nil, playerStyle)

			// Calculate next positions for collision detection
			newX1, newY1 := g.Player1.X, g.Player1.Y
			newX2, newY2 := g.Player2.X, g.Player2.Y

			switch g.Player1.Direction {
			case "up":
				newY1 -= 1
			case "down":
				newY1 += 1
			case "left":
				newX1 -= 2
			case "right":
				newX1 += 2
			}

			switch g.Player2.Direction {
			case "up":
				newY2 -= 1
			case "down":
				newY2 += 1
			case "left":
				newX2 -= 2
			case "right":
				newX2 += 2
			}

			g.Player1.X, g.Player1.Y = newX1, newY1
			g.Player2.X, g.Player2.Y = newX2, newY2

			g.Screen.Show()

			// Get screen dimensions for boundary checking
			width, height := g.Screen.Size()

			// Check collisions for Player 1
			player1Collision := false
			// Check boundaries
			if newX1 < 0 || newX1+1 >= width || newY1 < 0 || newY1 >= height {
				player1Collision = true
			}
			// Check trail collisions
			if !player1Collision && ((paintedMap[newX1] != nil && paintedMap[newX1][newY1]) ||
				(paintedMap[newX1+1] != nil && paintedMap[newX1+1][newY1])) {
				player1Collision = true
			}

			// Check collisions for Player 2
			player2Collision := false
			// Check boundaries
			if newX2 < 0 || newX2+1 >= width || newY2 < 0 || newY2 >= height {
				player2Collision = true
			}
			// Check trail collisions
			if !player2Collision && ((paintedMap[newX2] != nil && paintedMap[newX2][newY2]) ||
				(paintedMap[newX2+1] != nil && paintedMap[newX2+1][newY2])) {
				player2Collision = true
			}

			// Handle collisions
			if player1Collision && player2Collision {
				// Both players crashed - it's a tie
				// stop the game
				return
			} else if player1Collision {
				// Player 1 crashed - Player 2 wins
				//  stop the game
				return
			} else if player2Collision {
				// Player 2 crashed - Player 1 wins
				// stop the game
				return
			}
		}

	}

}

// func (g *Game) drawSprite(x int, y int, )
