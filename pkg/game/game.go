package game

import (
	"fmt"

	"github.com/carlostrejo2308/GoTakToe/pkg/board"
	"github.com/carlostrejo2308/GoTakToe/pkg/ia"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

// Game represents a game of tick-tack-toe
type Game struct {
	board      *board.Board
	Winner     piece.Player
	iaMechanic func(board.Board) (int, int)
}

// String returns a string representation of the game
func (g Game) String() string {
	// what a piece is represented by
	represent := map[piece.Player]string{
		piece.Empty: " ",
		piece.Ia:    "X",
		piece.Human: "O",
	}

	toPrint := ""

	// For each row, format the board
	for i, row := range g.board {
		toPrint += fmt.Sprintf("%s | %s | %s", represent[row[0]], represent[row[1]], represent[row[2]])
		if i < 2 {
			toPrint += "\n---------\n"
		}
	}

	return toPrint
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{
		board:      &board.Board{},
		Winner:     piece.Empty,
		iaMechanic: ia.Random,
	}
}

// isWinning checks if a player has won the game
func (g *Game) isWinning(player piece.Player) bool {

	// Check if the player has won in the board
	b := g.board
	w := b.IsWinning(player)

	// If the player has won, set the winner
	if w {
		g.Winner = player
	}

	return w
}

// StillPlaying returns true if the game has more turns to play
func (g *Game) StillPlaying() bool {
	b := g.board
	return !b.IsFull() && g.Winner == piece.Empty
}

func (g *Game) SetIa(ia func(board.Board) (int, int)) {
	g.iaMechanic = ia
}

// IaTurn plays a turn for the ia
func (g *Game) IaTurn() {

	b := g.board
	var x, y int
	played := false

	fmt.Println("I'm thinking...")
	for !played {
		x, y = g.iaMechanic(*b)

		if err := b.Play(piece.Ia, x, y); err != nil {
			continue
		} else {
			fmt.Printf("My turn: %d, %d\n", x, y)
			played = true
		}
	}
}

// Play calls the turn of both players and checks if someone has won
func (g *Game) Play() {
	b := g.board
	fmt.Println(g)

	// While the game is still playing
	for g.StillPlaying() {
		// Human turn
		b.HumanTurn()
		fmt.Println()
		fmt.Println(g)           // Print the board
		g.isWinning(piece.Human) // Check if the human has won
		if !g.StillPlaying() {
			break
		}

		// Ia turn
		g.IaTurn()
		fmt.Println()
		fmt.Println(g)
		g.isWinning(piece.Ia)
	}
}
