package main

import "fmt"

type Player int

const (
	empty Player = iota
	ia
	human
)

type Board [3][3]Player

type Game struct {
	board  *Board
	winner Player
}

func (g Game) String() string {
	represent := map[Player]string{
		empty: " ",
		ia:    "X",
		human: "O",
	}

	toPrint := ""

	for i, row := range g.board {
		toPrint += fmt.Sprintf("%s | %s | %s", represent[row[0]], represent[row[1]], represent[row[2]])
		if i < 2 {
			toPrint += "\n---------\n"
		}
	}

	return toPrint
}

func NewGame() *Game {
	return &Game{
		board:  &Board{},
		winner: empty,
	}
}

func (b *Board) isFull() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == empty {
				return false
			}
		}
	}

	return true
}

func (b *Board) isWinning(player Player) bool {
	// Check rows
	for _, row := range b {
		if row[0] == player && row[1] == player && row[2] == player {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return true
		}
	}

	// Check diagonals
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true
	}
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true
	}

	return false
}

func (b *Board) play(player Player, x, y int) error {
	if b[x][y] != empty {
		return fmt.Errorf("Cell (%d, %d) is already taken", x, y)
	}

	b[x][y] = player
	return nil
}

func (g *Game) StillPlaying() bool {
	b := g.board
	return !b.isFull() && g.winner == empty
}

func main() {
	g := NewGame()

	/* for b.StillPlaying() {
		b.HumanTurn()
		b.IATurn()
	} */

	fmt.Println(g)
}
