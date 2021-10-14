package ia

import (
	"math/rand"
	"time"

	"github.com/carlostrejo2308/GoTakToe/pkg/board"
	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

// iaRandom returns a random position of the board,
// it is used as a default ia
func Random(b board.Board) (int, int) {
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
func Challenger(b board.Board) (int, int) {
	var x, y int

	ogBoard := b

	// Check if there's a winning move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ogBoard.Play(piece.Ia, i, j)
			if win := ogBoard.IsWinning(piece.Ia); win {
				return i, j
			}
			ogBoard = b
		}
	}

	// Check if there is a losing move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ogBoard.Play(piece.Human, i, j)
			if win := ogBoard.IsWinning(piece.Human); win {
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
