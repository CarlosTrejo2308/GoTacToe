package board

import (
	"fmt"

	"github.com/carlostrejo2308/GoTakToe/pkg/piece"
)

type Board [3][3]piece.Player

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

func (b *Board) HumanTurn() {
	var x, y int
	played := false

	fmt.Print("Your turn (x, y): ")
	for !played {
		fmt.Scanf("%d %d", &x, &y)

		if err := b.Play(piece.Human, x, y); err != nil {
			fmt.Println(err)
		} else {
			played = true
		}
	}
}
