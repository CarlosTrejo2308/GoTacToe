package main

import (
	"fmt"
	"math/rand"
)

type Player int

const (
	empty Player = iota
	ia
	human
)

type Board [3][3]Player

type Game struct {
	board      *Board
	winner     Player
	iaMechanic func() (int, int)
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
		board:      &Board{},
		winner:     empty,
		iaMechanic: iaRandom,
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

func (g *Game) isWinning(player Player) bool {
	b := g.board
	w := b.isWinning(player)

	if w {
		g.winner = player
	}

	return w
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
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return fmt.Errorf("invalid coordinates")
	}
	if b[x][y] != empty {
		return fmt.Errorf("cell (%d, %d) is already taken", x, y)
	}

	b[x][y] = player
	return nil
}

func (g *Game) StillPlaying() bool {
	b := g.board
	return !b.isFull() && g.winner == empty
}

func (b *Board) HumanTurn() {
	var x, y int
	played := false

	fmt.Print("Your turn (x, y): ")
	for !played {
		fmt.Scanf("%d %d", &x, &y)

		if err := b.play(human, x, y); err != nil {
			fmt.Println(err)
		} else {
			played = true
		}
	}
}

func (g *Game) IaTurn() {

	b := g.board
	var x, y int
	played := false

	fmt.Println("I'm thinking...")
	for !played {
		x, y = g.iaMechanic()

		if err := b.play(ia, x, y); err != nil {
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
		g.isWinning(human)
		if !g.StillPlaying() {
			break
		}

		g.IaTurn()
		fmt.Println(g)
		g.isWinning(ia)
	}
}

func main() {
	g := NewGame()
	g.Play()

	switch g.winner {
	case empty:
		fmt.Println("It's a draw!")
	case human:
		fmt.Println("You win!")
	case ia:
		fmt.Println("I win!")
	}
}
