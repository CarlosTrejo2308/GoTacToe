package board

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

// Board is a 3x3 board
type Board [3][3]piece.Player

// IsFull returns true if the board is full
func (b *Board) IsFull() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == piece.Empty {
				return false
			}
		}
	}

	return true
}

// IsWinner returns true if the player has won
func (b *Board) IsWinning(player piece.Player) bool {
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

// Play places a piece in the board
func (b *Board) Play(player piece.Player, x, y int) error {
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return fmt.Errorf("invalid coordinates")
	}
	if b[x][y] != piece.Empty {
		return fmt.Errorf("cell (%d, %d) is already taken", x, y)
	}

	b[x][y] = player
	return nil
}

// HumanTurn gets the coordinates of the player's move and plays it
func (b *Board) HumanTurn() {
	var x, y int
	played := false

	for !played {
		fmt.Print("Your turn (x, y): ")
		x, y = Challenger(*b)
		//fmt.Scanf("%d %d", &x, &y)

		if err := b.Play(piece.Human, x, y); err != nil {
			fmt.Println(err)
		} else {
			played = true
		}
	}
}

func (b *Board) IsValidMove(x, y int) bool {
	return x >= 0 && x <= 2 && y >= 0 && y <= 2 && b[x][y] == piece.Empty
}

func Random(b Board) (int, int) {
	var x, y int

	rand.Seed(time.Now().Unix())

	x = rand.Intn(3)
	y = rand.Intn(3)

	wait := time.Duration(rand.Intn(700) + 100)
	time.Sleep(wait * time.Millisecond)

	return x, y
}

// Challenger is an IA that makes a move in the board
// Considering if he has a winning move or
// if the other player has a winning move, it blocks it
// or makes a random move
func Challenger(b Board) (int, int) {
	var x, y int

	ogBoard := b

	// Check if there's a winning move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ogBoard.Play(piece.Human, i, j)
			if win := ogBoard.IsWinning(piece.Human); win {
				return i, j
			}
			ogBoard = b
		}
	}

	// Check if there is a losing move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ogBoard.Play(piece.Ia, i, j)
			if win := ogBoard.IsWinning(piece.Ia); win {
				return i, j
			}
			ogBoard = b
		}
	}

	// Choose a random position
	valid := false
	for !valid {
		x, y = Random(b)
		if err := b.Play(piece.Ia, x, y); err != nil {
			continue
		} else {
			valid = true
		}
	}

	return x, y

}
