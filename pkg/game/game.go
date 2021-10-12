package game

import (
	"fmt"
	"math/rand"

	"github.com/carlostrejo2308/GoTakToe/pkg/board"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

type Game struct {
	board      *board.Board
	Winner     piece.Player
	iaMechanic func() (int, int)
}

func (g Game) String() string {
	represent := map[piece.Player]string{
		piece.Empty: " ",
		piece.Ia:    "X",
		piece.Human: "O",
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
		board:      &board.Board{},
		Winner:     piece.Empty,
		iaMechanic: iaRandom,
	}
}

func (g *Game) isWinning(player piece.Player) bool {
	b := g.board
	w := b.IsWinning(player)

	if w {
		g.Winner = player
	}

	return w
}

func (g *Game) StillPlaying() bool {
	b := g.board
	return !b.IsFull() && g.Winner == piece.Empty
}

func (g *Game) IaTurn() {

	b := g.board
	var x, y int
	played := false

	fmt.Println("I'm thinking...")
	for !played {
		x, y = g.iaMechanic()

		if err := b.Play(piece.Ia, x, y); err != nil {
			continue
		} else {
			fmt.Printf("My turn: %d, %d\n", x, y)
			played = true
		}
	}
}

func iaRandom() (int, int) {
	var x, y int

	x = rand.Intn(3)
	y = rand.Intn(3)

	return x, y
}

func (g *Game) Play() {
	b := g.board
	fmt.Println(g)

	for g.StillPlaying() {
		b.HumanTurn()
		fmt.Println(g)
		g.isWinning(piece.Human)
		if !g.StillPlaying() {
			break
		}

		g.IaTurn()
		fmt.Println(g)
		g.isWinning(piece.Ia)
	}
}
