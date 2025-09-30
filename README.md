# T-Tron

A Tron-style game for the terminal built with Go and tcell v2.

## Description

T-Tron is a classic Tron light cycle game that runs directly in your terminal. Players control light cycles that leave trails behind them, and the objective is to avoid crashing into walls or trails while trying to make opponents crash into yours.

This is currently a work in progress as I learn Go and tcell!

## Technology

- **Language**: Go 1.25.1+
- **UI Library**: [tcell v2](https://github.com/gdamore/tcell) for terminal-based graphics and input handling

## Requirements

- Go 1.25.1 or higher
- Terminal that supports color and cursor positioning

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/CJMeek/t-tron.git
   cd t-tron
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the game:
   ```bash
   go run .
   ```

## Game Controls

- Player 1: Arrow-keys
- Player 2: W A S D

## Work in Progress

This project is currently under active development. The following features are planned or being worked on:

### To-Do List

- [ ] **Collision Detection** - Implement proper collision detection for players and walls
- [ ] **Terminal Size Handling** - Fix terminal resizing and ensure proper game area scaling
- [ ] **Local Multiplayer Controls** - Fix and improve local multiplayer control scheme
- [ ] **Online Multiplayer** - Add network multiplayer functionality
- [ ] **GUI and Game Settings** - Add Bubbletea GUI to start the game and change settings about the game (e.g player speed, crashing into walls, game bounds(?))