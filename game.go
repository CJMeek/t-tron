package main

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

// TickDuration defines ~60fps update interval.
const TickDuration = 33 * time.Millisecond

// Winner codes: 0 none, 1 player1, 2 player2, 3 tie.
type Game struct {
	Screen  tcell.Screen
	Player1 Player
	Player2 Player
	painted map[int]map[int]bool
	Winner  int
	mu      sync.RWMutex // protects player positions and directions
}

// NewGame constructs a Game with internal trail map.
func NewGame(screen tcell.Screen, p1, p2 Player) *Game {
	return &Game{Screen: screen, Player1: p1, Player2: p2, painted: make(map[int]map[int]bool)}
}

// markTrail records a 2-cell wide horizontal player sprite trail.
func (g *Game) markTrail(x, y int, colour tcell.Color) {
	style := tcell.StyleDefault.Background(colour).Foreground(colour)
	g.Screen.SetContent(x, y, ' ', nil, style)
	g.Screen.SetContent(x+1, y, ' ', nil, style)
	if g.painted[x] == nil {
		g.painted[x] = make(map[int]bool)
	}
	g.painted[x][y] = true
	if g.painted[x+1] == nil {
		g.painted[x+1] = make(map[int]bool)
	}
	g.painted[x+1][y] = true
}

// SetDirection safely updates a player's direction
func (g *Game) SetPlayer1Direction(dir Direction) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Player1.SetDirection(dir)
}

func (g *Game) SetPlayer2Direction(dir Direction) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Player2.SetDirection(dir)
}

// Step advances the game state by one tick: move, paint, collision detect.
// Returns true if the game should continue.
func (g *Game) Step() bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.Winner != 0 { // already finished
		return false
	}

	// Compute next positions
	newX1, newY1 := g.Player1.X, g.Player1.Y
	switch g.Player1.Direction {
	case DirectionUp:
		newY1 -= 1
	case DirectionDown:
		newY1 += 1
	case DirectionLeft:
		newX1 -= PlayerWidth
	case DirectionRight:
		newX1 += PlayerWidth
	}

	newX2, newY2 := g.Player2.X, g.Player2.Y
	switch g.Player2.Direction {
	case DirectionUp:
		newY2 -= 1
	case DirectionDown:
		newY2 += 1
	case DirectionLeft:
		newX2 -= PlayerWidth
	case DirectionRight:
		newX2 += PlayerWidth
	}

	// Collision detection before committing positions
	width, height := g.Screen.Size()

	// Helper closure
	collides := func(x, y int) bool {
		if x < 0 || x+1 >= width || y < 0 || y >= height {
			return true
		}
		if (g.painted[x] != nil && g.painted[x][y]) || (g.painted[x+1] != nil && g.painted[x+1][y]) {
			return true
		}
		return false
	}

	p1Crash := collides(newX1, newY1)
	p2Crash := collides(newX2, newY2)

	// Decide winner
	if p1Crash && p2Crash {
		g.Winner = 3
	} else if p1Crash {
		g.Winner = 2
	} else if p2Crash {
		g.Winner = 1
	}

	if g.Winner != 0 { // game ends, do not paint future trails
		g.Screen.Show()
		return false
	}

	// Paint trails from previous positions
	g.markTrail(g.Player1.X, g.Player1.Y, g.Player1.Colour)
	g.markTrail(g.Player2.X, g.Player2.Y, g.Player2.Colour)

	// Commit movement
	g.Player1.X, g.Player1.Y = newX1, newY1
	g.Player2.X, g.Player2.Y = newX2, newY2

	// Draw white player sprites at NEW positions (after painting trails)
	playerStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	g.Screen.SetContent(g.Player1.X, g.Player1.Y, ' ', nil, playerStyle)
	g.Screen.SetContent(g.Player1.X+1, g.Player1.Y, ' ', nil, playerStyle)
	g.Screen.SetContent(g.Player2.X, g.Player2.Y, ' ', nil, playerStyle)
	g.Screen.SetContent(g.Player2.X+1, g.Player2.Y, ' ', nil, playerStyle)

	g.Screen.Show()
	return true
}

// Run drives the game loop until quit is signalled or a winner is decided.
func (g *Game) Run(quit <-chan struct{}) {
	ticker := time.NewTicker(TickDuration)
	defer ticker.Stop()
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			if !g.Step() { // finished
				return
			}
		}
	}
}
